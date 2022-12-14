import random

# WORD BANKS FOR CODE COMBINATIONS
first_word_bank = [
  'PEANUT',
  'EXIT',
  'TUNE',
  'BONFIRE',
  'JAW',
  'MOTHER',
  'FIELD', # 2.0
]

second_word_bank = [
  'ASTRONAUT',
  'MONSTERS',
  'SOUR',
  'BLEND',
  'PLANE',
  'PLANTAIN', # 2.0
  'NUKE', # 2.0
]

third_word_bank = [
  'COCONUT',
  'SUIT',
  'BLUE',
  'ROCKY',
  'MOON',
  'TREETOP',
  'RASPBERRY', # 2.0
]

fourth_word_bank = [
  'FIELD',
  'PLANE',
  'CLOUD',
  'NUKE',
  'SHOWER',
  'BUGS',
  'BEAN', # 2.0
  'PEEK', # 2.0
]

# FUNCTION TO PERMEATE EVERY POSSIBLE COMBINATION
def get_code_options():
  options = []

  for first_word in first_word_bank:
    for second_word in second_word_bank:
      for third_word in third_word_bank:
        for fourth_word in fourth_word_bank:
          options.append(f'{first_word} {second_word} {third_word} {fourth_word}')

  return options

# GET OPTIONS AND SHUFFLE LIST
code_options = get_code_options()
random.shuffle(code_options)

# PRINT EACH CODE OPTION LINE BY LINE
for code in code_options:
  print(code)
