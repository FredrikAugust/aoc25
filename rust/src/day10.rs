use rayon::iter::{IntoParallelRefIterator, ParallelIterator};
use regex::Regex;
use std::{collections::HashMap, fs::File, io::Read, ops::Index};
use z3::{Optimize, ast::Int};

#[derive(Debug)]
struct Machine {
    buttons: Vec<Vec<usize>>,
    joltages: Vec<i32>,
}

pub fn day10part2() -> u64 {
    let mut buf: String = Default::default();
    File::open("src/10-input.txt")
        .unwrap()
        .read_to_string(&mut buf)
        .unwrap();

    let pattern = Regex::new(r"\[([.#]+)\] ((\([\d,]+\) ?)+) \{([\d,]+)\}").unwrap();
    let machines = buf
        .lines()
        .map(|line| {
            let parts = pattern.captures(line).unwrap();
            let buttons = parts.get(2).unwrap();
            let joltages = parts.get(4).unwrap();

            Machine {
                buttons: buttons
                    .as_str()
                    .split_whitespace()
                    .map(|group| {
                        let parenthesis_content = group.trim_matches(['(', ')']);
                        parenthesis_content
                            .split(',')
                            .map(|s| s.parse().unwrap())
                            .collect()
                    })
                    .collect(),
                joltages: joltages
                    .as_str()
                    .split(',')
                    .map(|s| s.parse().unwrap())
                    .collect(),
            }
        })
        .collect::<Vec<_>>();

    machines.par_iter().map(solve_machine).sum()
}

fn solve_machine(machine: &Machine) -> u64 {
    let mut map: HashMap<usize, Vec<usize>> = Default::default();
    let inputs = machine
        .buttons
        .iter()
        .enumerate()
        .map(|(i, b)| {
            let value = Int::fresh_const(format!("{:?}", b).as_str());

            for button in b.iter() {
                match map.get_mut(button) {
                    Some(vec) => vec.push(i),
                    None => {
                        map.insert(*button, vec![i]);
                    }
                }
            }

            value
        })
        .collect::<Vec<_>>();

    let solver = Optimize::new();

    for input in inputs.iter() {
        solver.assert(&input.ge(0));
    }

    for (i, joltage) in machine.joltages.iter().enumerate() {
        let relevant_inputs = map
            .get(&i)
            .unwrap()
            .into_iter()
            .map(|input| inputs.index(*input))
            .collect::<Vec<_>>();

        for relevant_input in &relevant_inputs {
            solver.assert(&relevant_input.le(*joltage));
        }
        solver.assert(&Int::add(relevant_inputs.as_slice()).eq(*joltage));
    }

    solver.minimize(&Int::add(inputs.as_slice()));
    solver.check(&[]);
    let model = solver.get_model().unwrap();

    inputs
        .iter()
        .map(|i| model.eval(i, true).unwrap().as_u64().unwrap())
        .sum::<u64>()
}
