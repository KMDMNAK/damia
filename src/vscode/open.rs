// start up -d container using service name and attach to vscode.

extern crate yaml_rust;

use std::{fs, string};
use std::ops::Index;
use yaml_rust::{YamlEmitter, YamlLoader};

fn read(path: &str) -> string::String {
    let contents = fs::read_to_string(path);
    println!("{:?}", contents);
    return contents.unwrap();
}

pub fn get_yml(path: &str) {
    // let contents = fs::read_to_string(path).expect("Something went wrong reading the file");
    let contents = read(path);
    let docs = YamlLoader::load_from_str(&contents).unwrap();
    // Multi document support, doc is a yaml::Yaml
    let doc = &docs[0];

    // Debug support
    println!("{:?}", doc);

    // Dump the YAML object
    let mut out_str = String::new();
    {
        let mut emitter = YamlEmitter::new(&mut out_str);
        emitter.dump(doc).unwrap(); // dump the YAML object to a String
    }
    let services = doc.index("services");
    println!("{:?}", services);
}
