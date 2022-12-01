use std::{io::{self, BufWriter, Stdin}, thread::current, i32::MAX};

fn main() {
    let stdin = io::stdin();
    let lines = stdin.lines();

    let mut all_elves = Vec::new();
    let mut current_elf = Vec::new();

    for line in lines {
        let text = line.unwrap();

        if text.is_empty() {
            all_elves.push(current_elf.clone());
            println!("[{}]", current_elf.iter().map( |id: &i32| id.to_string() + ",").collect::<String>());
            current_elf.truncate(0);
        } else {
            let calories = text.parse::<i32>().unwrap();
            println!("{}", calories);
            current_elf.push(calories);
        }

    }

    let mut max_count = 0;
    let mut counts = Vec::new();
    for elf in all_elves {
        let sum = elf.iter().sum::<i32>();
        max_count = std::cmp::max(max_count, sum);
        counts.push(sum);
    }

    println!("{}", max_count);

    counts.sort();
    counts.reverse();
    println!("{}", counts[0] + counts[1] + counts[2]);

}
