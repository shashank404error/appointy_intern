package accounty

import "go.uber.org/fx"

// Module is the fx module encapsulating all the providers of the package 
var Module = fx.Provide(
	NewquickbookServerServer,
)