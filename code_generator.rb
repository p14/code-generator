def get_code_options
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

  options = []

  first_word_bank.each do |first_word|
    second_word_bank.each do |second_word|
      third_word_bank.each do |third_word|
        fourth_word_bank.each do |fourth_word|
          options << "#{first_word} #{second_word} #{third_word} #{fourth_word}"
        end
      end
    end
  end
  return options
end

code_options = get_code_options()
shuffled_options = code_options.shuffle

shuffled_options.each do |code|
  puts code
end
