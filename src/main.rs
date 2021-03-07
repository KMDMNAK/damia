extern crate clap;
use clap::{App, Arg, SubCommand};

use damia::command::generate;

fn main(){
    let shell_arg = Arg::with_name("shell")
        .value_name("SHELL")
        .help(
            "The name of the currently running shell\nCurrently supported options: bash, zsh, fish, powershell, ion, elvish",
        )
        .required(false);

    let app = App::new("clapex")
        .version("0.1.0")
        .author("myname <myname@mail.com>")
        .about("Clap Example CLI")

        .subcommand(
            SubCommand::with_name("generate")
                .about("generate templetes for each version")
                .arg(&shell_arg)
        )
        .subcommand(
            SubCommand::with_name("help")
                .about("Prints the shell function used to execute starship")
                .arg(&shell_arg)
        )
        .arg(Arg::with_name("arg1").short("1").long("argument1").takes_value(true))
        .arg(Arg::with_name("arg2").short("2").long("argument2").takes_value(true));

    let matches = app.clone().get_matches();
    match matches.subcommand() {
        ("generate", Some(sub_m)) => {
            // let shell_name = sub_m.value_of("shell").expect("Shell name missing.");
            generate::command()
        }
        ("help",Some(sub_m)) => {},
        (command, _) => println!("unreach")
    }
}