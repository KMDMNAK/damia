extern crate clap;
// extern crate command;

use clap::{App, Arg, SubCommand};

use damia::command::{generate, vscode};

fn main() {
    let shell_arg = Arg::with_name("shell")
        .value_name("SHELL")
        .help(
            "The name of the currently running shell\nCurrently supported options: bash, zsh, fish, powershell, ion, elvish",
        )
        .required(false);
    let vscode_open_arg = Arg::with_name("SERVICE")
        .value_name("SERVICE")
        .help("Service name to open container attached by vscode.")
        .required(true);

    let app = App::new("clapex")
        .version("0.1.0")
        .author("myname <myname@mail.com>")
        .about("Clap Example CLI")
        .subcommand(
            SubCommand::with_name("generate")
                .about("generate templetes for each version")
                .arg(&shell_arg),
        )
        .subcommand(
            SubCommand::with_name("help")
                .about("Prints the shell function used to execute starship")
                .arg(&shell_arg),
        )
        .subcommand(
            SubCommand::with_name("vscode")
                .about("manipulate vscode")
                .subcommand(SubCommand::with_name("open").arg(&vscode_open_arg)),
        )
        .arg(
            Arg::with_name("arg1")
                .short("1")
                .long("argument1")
                .takes_value(true),
        )
        .arg(
            Arg::with_name("arg2")
                .short("2")
                .long("argument2")
                .takes_value(true),
        );

    let matches = app.clone().get_matches();
    match matches.subcommand() {
        ("generate", Some(_sub_m)) => {
            // let shell_name = _sub_m.value_of("shell").expect("Shell name missing.");
            generate::command()
        }
        ("help", Some(_sub_m)) => {}
        ("vscode", Some(_sub_m)) => {
            match _sub_m.subcommand() {
                ("open", Some(sub_sub_m)) => {
                    println!("open!!");
                    let service = sub_sub_m.value_of("SERVICE");
                    match service {
                        None => println!("No idea what your favorite number is."),
                        Some(s) => {
                            println!("{}", s);
                            vscode::open::get_yml("docker-compose.yml");
                        }
                    }
                }
                (_command, _) => println!("unreach"),
            }
            println!("vscode!!");
        }
        (_command, _) => println!("unreach"),
    }
}
