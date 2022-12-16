#[derive(Clone)]
pub struct Rule {
    before: Vec<char>,
    after: Vec<char>,
    is_terminate: bool,
}

impl Rule {
    pub fn new(before: &str, after: &str, is_terminate: bool) -> Self {
        Rule {
            before: before.chars().collect(),
            after: after.chars().collect(),
            is_terminate,
        }
    }
}

pub struct Markov {
    rules: Vec<Rule>,
}

impl Markov {
    pub fn new(code: &str) -> Result<Self, &str> {
        match Markov::parse(code) {
            Ok(rules) => Ok(Markov { rules }),
            Err(msg) => Err(msg),
        }
    }

    fn parse(code: &str) -> Result<Vec<Rule>, &str> {
        let code_chars: Vec<char> = code.chars().collect();
        let mut rules = vec![];
        let mut i = 0 as usize;
        while i < code_chars.len() {
            let mut before = vec![];
            let mut after = vec![];
            let is_terminate;
            
            while i < code_chars.len() && code_chars[i] != ':' {
                before.push(code_chars[i]);
                i += 1;
            }

            if i < code_chars.len() && code_chars[i] == ':' {
                if i + 1 < code_chars.len() && code_chars[i+1] == ':' {
                    is_terminate = true;
                    i += 2;
                } else {
                    is_terminate = false;
                    i += 1;
                }   
            } else {
                return Err("Syntax Error");
            }

            while i < code_chars.len() && code_chars[i] != '\n' {
                after.push(code_chars[i]);
                i += 1;
            }

            rules.push(Rule { before, after, is_terminate });

            i += 1;
        }
        return Ok(rules);
    }

    pub fn compute(&self, input: &str) -> Vec<char> {
        let mut cur_text: Vec<char> = input.chars().collect();

        loop {
            let mut some_rule_match = false;
            for rule in self.rules.clone() {
                let mut pattern_match = true; // for matching an empty string

                if cur_text.len() < rule.before.len() {
                    continue;
                }
                for i in (0 as usize)..((cur_text.len() - rule.before.len() + 1) as usize) {
                    pattern_match = true;
                    for j in (0 as usize)..(rule.before.len() as usize) {
                        if cur_text[i + j] != rule.before[j] {
                            pattern_match = false;
                            break;
                        }
                    }

                    if pattern_match {//println!("ok");
                        some_rule_match = true;
                        let mut next_text: Vec<char> = vec![];
                        for j in 0..i {
                            next_text.push(cur_text[j]);
                        }
                        for c in rule.after {
                            next_text.push(c);
                        }
                        for j in i + rule.before.len()..cur_text.len() {
                            next_text.push(cur_text[j]);
                        }
                        cur_text = next_text;

                        if rule.is_terminate {
                            return cur_text;
                        }

                        break;
                    }
                }

                if pattern_match {
                    break;
                }
            }

            if !some_rule_match {
                break;
            }
        }

        return cur_text;
    }
}