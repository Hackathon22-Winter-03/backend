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
    text: Vec<char>,
}

impl Markov {
    pub fn new(code: &str) -> Result<Self, &str> {
        match Markov::parse(code) {
            Ok(rules) => Ok(Markov {
                rules: rules,
                text: Vec::<char>::new(),
            }),
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
                if i + 1 < code_chars.len() && code_chars[i + 1] == ':' {
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

            rules.push(Rule {
                before,
                after,
                is_terminate,
            });

            i += 1;
        }
        return Ok(rules);
    }

    pub fn set_text(&mut self, text: &str) {
        self.text = text.chars().collect();
    }

    pub fn run(&mut self) -> Vec<char> {
        loop {
            let (next_text, is_terminate, is_ended) = self.step();
            if is_terminate {
                return next_text;
            } else {
                self.text = next_text;
            }
        }
    }

    pub fn step(&mut self) -> (Vec<char>, bool, bool) {
        let mut flag: bool = false;
        for rule in self.rules.clone() {
            if self.text.len() < rule.before.len() {
                continue;
            }

            for i in (0 as usize)..((self.text.len() - rule.before.len() + 1) as usize) {
                let mut pattern_match = true;
                for j in (0 as usize)..(rule.before.len() as usize) {
                    if self.text[i + j] != rule.before[j] {
                        pattern_match = false;
                        break;
                    }
                }

                if pattern_match {
                    flag = true;
                    let mut next_text: Vec<char> = vec![];
                    for j in 0..i {
                        next_text.push(self.text[j]);
                    }
                    for c in rule.after {
                        next_text.push(c);
                    }
                    for j in i + rule.before.len()..self.text.len() {
                        next_text.push(self.text[j]);
                    }
                    return (next_text, rule.is_terminate);
                }
            }
        }

        return (self.text.clone(), true, !flag);
    }
}
