use criterion::{criterion_group, criterion_main, Criterion};
use rand_core::OsRng;

use penumbra_proof_setup::{
    log::{ContributionHash, Hashable},
    {Phase1CRSElements, Phase1Contribution, Phase1RawContribution},
};

fn run_phase1_prove(parent: ContributionHash, old: &Phase1CRSElements) -> Phase1Contribution {
    Phase1Contribution::make(&mut OsRng, parent, old)
}

fn run_phase1_verify(contribution: Phase1RawContribution, parent: &Phase1CRSElements) {
    let validated_contribution = contribution
        .validate(&mut OsRng)
        .expect("this is a valid contribution");
    let _ = validated_contribution.is_linked_to(parent);
}

fn phase1_benchmarks(c: &mut Criterion) {
    // Generate contribution for degree = 37,061
    // (size of largest proof)
    let d = 1 << 16;
    let root = Phase1CRSElements::root(d);
    let root_hash = root.hash();

    c.bench_function("phase 1 prove", |b| {
        b.iter(|| run_phase1_prove(root_hash, &root))
    });

    let new_contribution = Phase1Contribution::make(&mut OsRng, root_hash, &root);
    c.bench_function("phase 1 verify", |b| {
        b.iter(|| run_phase1_verify(new_contribution.clone().into(), &root))
    });
}

criterion_group! {
  name = benches;
  config = Criterion::default().sample_size(10);
  targets = phase1_benchmarks
}
criterion_main!(benches);