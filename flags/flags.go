package flags

import "github.com/urfave/cli/v2"

const evnVarPrefix = "MARKET"

func prefixEnvVars(name string) []string {
	return []string{evnVarPrefix + "_" + name}
}

var (
	MigrationsFlag = &cli.StringFlag{
		Name:    "migrations-dir",
		Value:   "./migrations",
		Usage:   "path for database migrations",
		EnvVars: prefixEnvVars("MIGRATIONS_DIR"),
	}

	// RpcHostFlag RPC Service
	RpcHostFlag = &cli.StringFlag{
		Name:     "rpc-host",
		Usage:    "The host of the rpc",
		EnvVars:  prefixEnvVars("RPC_HOST"),
		Required: true,
	}
	RpcPortFlag = &cli.IntFlag{
		Name:     "rpc-port",
		Usage:    "The port of the rpc",
		EnvVars:  prefixEnvVars("RPC_PORT"),
		Required: true,
	}

	// rest  RPC Service
	RestHostFlag = &cli.StringFlag{
		Name:     "rest-host",
		Usage:    "The host of the rest",
		EnvVars:  prefixEnvVars("REST_HOST"),
		Required: true,
	}
	RestPortFlag = &cli.IntFlag{
		Name:     "rest-port",
		Usage:    "The port of the rest",
		EnvVars:  prefixEnvVars("REST_PORT"),
		Required: true,
	}

	// Metrix RPC Service
	MetricsHostFlag = &cli.StringFlag{
		Name:     "rest-host",
		Usage:    "The host of the metrics",
		EnvVars:  prefixEnvVars("METRICS_HOST"),
		Required: true,
	}
	MetricsPortFlag = &cli.IntFlag{
		Name:     "rest-port",
		Usage:    "The port of the metrics",
		EnvVars:  prefixEnvVars("METRICS_PORT"),
		Required: true,
	}

	// MasterDb Flags
	MasterDbHostFlag = &cli.StringFlag{
		Name:     "master-db-host",
		Usage:    "The host of the master database",
		EnvVars:  prefixEnvVars("MASTER_DB_HOST"),
		Required: true,
	}
	MasterDbPortFlag = &cli.IntFlag{
		Name:     "master-db-port",
		Usage:    "The port of the master database",
		EnvVars:  prefixEnvVars("MASTER_DB_PORT"),
		Required: true,
	}
	MasterDbUserFlag = &cli.StringFlag{
		Name:     "master-db-user",
		Usage:    "The user of the master database",
		EnvVars:  prefixEnvVars("MASTER_DB_USER"),
		Required: true,
	}
	MasterDbPasswordFlag = &cli.StringFlag{
		Name:     "master-db-password",
		Usage:    "The host of the master database",
		EnvVars:  prefixEnvVars("MASTER_DB_PASSWORD"),
		Required: true,
	}
	MasterDbNameFlag = &cli.StringFlag{
		Name:     "master-db-name",
		Usage:    "The db name of the master database",
		EnvVars:  prefixEnvVars("MASTER_DB_NAME"),
		Required: true,
	}

	// Slave DB  flags
	SlaveDbHostFlag = &cli.StringFlag{
		Name:    "slave-db-host",
		Usage:   "The host of the slave database",
		EnvVars: prefixEnvVars("SLAVE_DB_HOST"),
	}
	SlaveDbPortFlag = &cli.IntFlag{
		Name:    "slave-db-port",
		Usage:   "The port of the slave database",
		EnvVars: prefixEnvVars("SLAVE_DB_PORT"),
	}
	SlaveDbUserFlag = &cli.StringFlag{
		Name:    "slave-db-user",
		Usage:   "The user of the slave database",
		EnvVars: prefixEnvVars("SLAVE_DB_USER"),
	}
	SlaveDbPasswordFlag = &cli.StringFlag{
		Name:    "slave-db-password",
		Usage:   "The host of the slave database",
		EnvVars: prefixEnvVars("SLAVE_DB_PASSWORD"),
	}
	SlaveDbNameFlag = &cli.StringFlag{
		Name:    "slave-db-name",
		Usage:   "The db name of the slave database",
		EnvVars: prefixEnvVars("SLAVE_DB_NAME"),
	}
)

var requireFlags = []cli.Flag{
	MigrationsFlag,
	RpcHostFlag,
	RpcPortFlag,
	RestHostFlag,
	RestPortFlag,

	MasterDbHostFlag,
	MasterDbPortFlag,
	MasterDbUserFlag,
	MasterDbPasswordFlag,
	MasterDbNameFlag,
}

var optionalFlags = []cli.Flag{
	SlaveDbHostFlag,
	SlaveDbPortFlag,
	SlaveDbUserFlag,
	SlaveDbPasswordFlag,
	SlaveDbNameFlag,
}

func init() {
	Flags = append(requireFlags, optionalFlags...)
}

var Flags []cli.Flag
