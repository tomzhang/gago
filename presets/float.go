package presets

import "github.com/MaxHalford/gago"

// Float returns a configuration for minimizing continuous mathematical
// functions with a given number of variables.
func Float(n int, function func([]float64) float64) gago.GA {
	return gago.GA{
		NbPopulations: 2,
		NbIndividuals: 30,
		NbGenes:       n,
		Ff: gago.Float64Function{
			Image: function,
		},
		Initializer: gago.InitUniformF{
			Lower: -1,
			Upper: 1,
		},
		Model: gago.ModGenerational{
			Selector: gago.SelTournament{
				NbParticipants: 3,
			},
			Crossover: gago.CrossUniformF{},
			Mutator: gago.MutNormalF{
				Rate: 0.5,
				Std:  3,
			},
			MutRate: 0.5,
		},
		Migrator:     gago.MigShuffle{},
		MigFrequency: 10,
	}
}
