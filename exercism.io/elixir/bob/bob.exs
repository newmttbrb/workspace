defmodule Teenager do
  def hey(input) do
    cond do
    	ask_question(input) -> "Sure."
      silence(input) -> "Fine. Be that way!"
    	yell(input) -> "Woah, chill out!"
    	true -> "Whatever."
    end
  end

  defp ask_question(input), do: String.ends_with?(input,"?")

  defp silence(input), do: String.strip(input) == ""

  defp yell(input) do
    letters_only = remove_numbers(remove_punctuation(input))
    words = String.split(letters_only)
    Enum.find(words,&loud_word/1)
  end	

  defp remove_punctuation(word), do: String.replace(word,~r/[\/\~\`\!\@\#\$\%\^\&\*\(\)\-\_\=\+\{\}\[\]\|\"\'\:\;\,\.\<\>\?\\]/,"")

  defp remove_numbers(word), do: String.replace(word,~r/[0123456789]/,"")

  defp loud_word(word), do: word == String.upcase(word)

end
