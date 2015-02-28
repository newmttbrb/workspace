defmodule Words do
  @doc """
  Count the number of words in the sentence.

  Words are compared case-insensitively.
  """
  @spec count(String.t) :: HashDict.t
  def count(sentence) do
  	words = divide_into_words(underscores_are_spaces(remove_punctuation(sentence)))
	count_words(HashDict.new(),words)  
  end

  defp count_words(dict,[]), do: dict
  defp count_words(dict,[h|t]) do
  	lower = String.downcase(h)
    v = Dict.get(dict,lower,0)
    d = Dict.put(dict,lower,v+1)
    count_words(d,t)	  	
  end

  defp divide_into_words(sentence), do: String.split(sentence)

  defp remove_punctuation(sentence) do
     String.replace(sentence,~r/[\!\"\'\:\;\,\.\?\@\#\$\%\^\&\*\(\)]/,"")
  end

  defp underscores_are_spaces(sentence), do: String.replace(sentence,~r/[_]/," ")

end
