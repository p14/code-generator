import random

# Word banks for code combinations
first_word_bank = [
    'PEANUT',
    'EXIT',
    'TUNE',
    'BONFIRE',
    'JAW',
    'MOTHER',
]

second_word_bank = [
    'ASTRONAUT',
    'MONSTERS',
    'SOUR',
    'BLEND',
    'PLANE',
]

third_word_bank = [
    'COCONUT',
    'SUIT',
    'BLUE',
    'ROCKY',
    'MOON',
    'TREETOP',
]

fourth_word_bank = [
    'FIELD',
    'PLANE',
    'CLOUD',
    'NUKE',
    'SHOWER',
    'BUGS',
]

# Function to permeate every possible combination
def get_code_options():
    options = []

    for first_word in first_word_bank:
        for second_word in second_word_bank:
            for third_word in third_word_bank:
                for fourth_word in fourth_word_bank:
                    options.append(f'{first_word} {second_word} {third_word} {fourth_word}')

    return options

# Get options and shuffle the list
code_options = get_code_options()
random.shuffle(code_options)

# Print each code option line-by-line
for code in code_options:
    print(code)
