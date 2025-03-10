package cli

import (
    "os"
	"github.com/pierrestoffe/tulip/pkg/app"
	"github.com/pierrestoffe/tulip/pkg/proxy"
	"github.com/pierrestoffe/tulip/pkg/util/log"
)

func Execute() error {
    if len(os.Args) < 2 {
        log.PrintInfo("Usage: tulip <command> [options]")
		log.PrintInfo("Run 'tulip help' for usage information.")
        os.Exit(1)
    }

    command := os.Args[1]
    args := os.Args[2:]

    switch command {
    case "help":
        log.PrintInfo("to do")
    case "init":
        app.Initialize()
    case "proxy":
        help := "Usage: tulip proxy {start|stop|restart}"
        if len(args) == 0 {
            log.PrintInfo(help)
            os.Exit(1)
        } else {
            switch args[0] {
            case "help":
                log.PrintInfo(help)
                os.Exit(1)
            case "start":
                proxy.Start()
            case "stop":
                proxy.Stop()
            case "restart":
                proxy.Restart()
            default:
                log.PrintError("Error: Unknown command "+ args[0])
                log.PrintInfo(help)
                os.Exit(1)
            }
        }
    // case "start", "up":
    //     startProject()
    // case "stop", "down":
    //     stopProject()
    // case "restart":
    //     if len(args) > 0 && args[0] == "proxy" {
    //         fmt.Println("Restarting Traefik...")
    //         if err := exec.Command("docker-compose", "restart").Run(); err != nil {
    //             fmt.Println("Error restarting Traefik:", err)
    //         }
    //         fmt.Println("Traefik restarted.")
    //     } else {
    //         restartProject()
    //     }
    // case "create", "new":
    //     if len(args) < 1 {
    //         fmt.Println("Usage: tulip create <project-name> [project-type] [domain]")
    //         os.Exit(1)
    //     }
    //     projectName := args[0]
    //     projectType := "craft"
    //     domain := projectName + ".loc"
    //     if len(args) > 1 {
    //         projectType = args[1]
    //     }
    //     if len(args) > 2 {
    //         domain = args[2]
    //     }
    //     createProject(projectName, projectType, domain)
    // case "ssh", "shell":
    //     service := "php"
    //     if len(args) > 0 {
    //         service = args[0]
    //     }
    //     sshContainer(service)
    // case "exec":
    //     if len(args) < 2 {
    //         fmt.Println("Usage: tulip exec <service> <command>")
    //         os.Exit(1)
    //     }
    //     service := args[0]
    //     command := args[1]
    //     execCommand(service, command)
    // case "composer":
    //     composerCommand(args...)
    // case "craft":
    //     craftCommand(args...)
    // case "npm":
    //     npmCommand(args...)
    // case "yarn":
    //     yarnCommand(args...)
    // case "domain":
    //     if len(args) < 2 {
    //         fmt.Println("Usage: tulip domain {add|remove|list} [domain]")
    //         os.Exit(1)
    //     }
    //     action := args[0]
    //     domain := args[1]
    //     domainCommand(action, domain)
    // case "db":
    //     if len(args) < 1 {
    //         fmt.Println("Usage: tulip db {export|import|connect|tableplus|sequelpro} [filename]")
    //         os.Exit(1)
    //     }
    //     action := args[0]
    //     filename := ""
    //     if len(args) > 1 {
    //         filename = args[1]
    //     }
    //     dbCommand(action, filename)
    // case "logs":
    //     service := ""
    //     follow := "true"
    //     if len(args) > 0 {
    //         service = args[0]
    //     }
    //     if len(args) > 1 {
    //         follow = args[1]
    //     }
    //     logsCommand(service, follow)
    // case "status", "ps":
    //     statusCommand()
    // case "version":
    //     fmt.Println("Tulip version", version)
    // case "help", "--help", "-h":
    //     fmt.Println("Tulip - A development environment manager")
    //     fmt.Println("")
    //     fmt.Println("Usage: tulip <command> [options]")
    //     fmt.Println("")
    //     fmt.Println("Commands:")
    //     fmt.Println("  init                 Initialize Tulip")
    //     fmt.Println("  start|up [proxy]     Start project or Traefik proxy")
    //     fmt.Println("  stop|down [proxy]    Stop project or Traefik proxy")
    //     fmt.Println("  restart [proxy]      Restart project or Traefik proxy")
    //     fmt.Println("  create|new <name>    Create a new project")
    //     fmt.Println("  ssh|shell [service]  SSH into a container (default: php)")
    //     fmt.Println("  exec <service> <cmd> Execute a command in a container")
    //     fmt.Println("  composer <cmd>       Run Composer commands")
    //     fmt.Println("  craft <cmd>          Run Craft CMS commands")
    //     fmt.Println("  npm <cmd>            Run npm commands")
    //     fmt.Println("  yarn <cmd>           Run Yarn commands")
    //     fmt.Println("  domain <action>      Manage domains (add, remove, list)")
    //     fmt.Println("  db <action>          Database operations (export, import, connect)")
    //     fmt.Println("  logs [service]       View container logs")
    //     fmt.Println("  status|ps            Show project status")
    //     fmt.Println("  version              Show version information")
    //     fmt.Println("  help                 Show this help message")
    //     fmt.Println("")
    //     fmt.Println("Examples:")
    //     fmt.Println("  tulip init                      Initialize Tulip")
    //     fmt.Println("  tulip create myproject          Create a new project")
    //     fmt.Println("  tulip start                     Start the current project")
    //     fmt.Println("  tulip start proxy               Start the Traefik proxy")
    //     fmt.Println("  tulip ssh                       SSH into the PHP container")
    //     fmt.Println("  tulip exec db mysql -u root -p  Execute a command in the DB container")
    //     fmt.Println("  tulip composer install          Run Composer install")
    //     fmt.Println("  tulip craft migrate/all         Run Craft migrations")
    //     fmt.Println("  tulip domain add api.mysite.loc Add a domain to the project")
    //     fmt.Println("  tulip db export backup.sql      Export the database")
    //     fmt.Println("  tulip logs web                  View logs for the web container")
    default:
        log.PrintError("Error: Unknown command "+ command)
        log.PrintInfo("Run 'tulip help' for usage information.")
        os.Exit(1)
    }

    return nil
}
