package gago

import "math/rand"

// Mutator modifies an individual by replacing it's genes with new values.
type Mutator interface {
	Apply(indi *Individual, generator *rand.Rand)
}

// MutNormalF modifies a float gene if a coin toss is under a defined mutation
// ate. It does so for each gene. The new gene value is a random value sampled
// from a normal distribution centered on the gene's current value and with the
// intensity parameter as it's standard deviation. Only works for floating point
// values.
type MutNormalF struct {
	Rate float64 // Mutation rate for each gene
	Std  float64 // Standard deviation
}

// Apply normal mutation.
func (mut MutNormalF) Apply(indi *Individual, generator *rand.Rand) {
	for i := range indi.Genome {
		// Flip a coin and decide to mutate or not
		if generator.Float64() < mut.Rate {
			// Sample from a normal distribution
			indi.Genome[i] = indi.Genome[i].(float64) * generator.NormFloat64() * mut.Std
		}
	}
}

// MutSplice splices a genome in 3 and glues the parts back together in another
// order.
type MutSplice struct{}

// Apply splice mutation.
func (mut MutSplice) Apply(indi *Individual, generator *rand.Rand) {
	// Choose where to start and end the splice
	var (
		end   = rand.Int() % len(indi.Genome)
		start = rand.Int() % end
	)
	// Split the genome into two
	var inner = make(Genome, end-start)
	copy(inner, indi.Genome[start:end])
	var outer = append(indi.Genome[:start], indi.Genome[end:]...)
	// Choose where to insert the splice
	var insert = rand.Int() % len(outer)
	// Splice and insert
	indi.Genome = append(
		outer[:insert],
		append(inner, outer[insert:]...)...,
	)
}

// MutPermute permutes two genes.
type MutPermute struct {
	// Maximum number of permutation
	Max int
}

// Apply permutation mutation.
func (mut MutPermute) Apply(indi *Individual, generator *rand.Rand) {
	for i := 0; i <= generator.Intn(mut.Max); i++ {
		// Choose two points on the genome
		var (
			points = generator.Perm(len(indi.Genome))[:2]
			i      = points[0]
			j      = points[1]
		)
		// Permute the genes
		indi.Genome[i], indi.Genome[j] = indi.Genome[j], indi.Genome[i]
	}
}

// MutUniformS permutes two genes.
type MutUniformS struct {
	Corpus []string // Corpus to replace genes with
}

// Apply permutation mutation.
func (mut MutUniformS) Apply(indi *Individual, generator *rand.Rand) {
	// Choose a random element from the corpus
	var element = mut.Corpus[generator.Intn(len(mut.Corpus))]
	// Choose a position on the individual's genome
	var p = generator.Intn(len(indi.Genome))
	// Replace the gene at the chosen position with the chosen element
	indi.Genome[p] = element
}
