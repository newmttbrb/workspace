defmodule ListOps do
  # Please don't use any external modules (especially List) in your
  # implementation. The point of this exercise is to create these basic functions
  # yourself.
  # 
  # Note that `++` is a function from an external module (Kernel, which is
  # automatically important`) and so shouldn't be used either.
 
  @spec count(list) :: non_neg_integer
  def  count(l), do: reduce(l,0,fn(_,acc) -> acc + 1 end)

  @spec reverse(list) :: list
  def  reverse(l), do: reduce(l,[],fn(h,acc)-> [h|acc] end)

  @spec map(list, (any -> any)) :: list
  def map(l, f), do: reduce(reverse(l),[],fn(h,acc) -> [f.(h)|acc] end)

  @spec filter(list, (any -> as_boolean(term))) :: list
  def filter(l, f), do: reverse l  |> reduce([],fn(h,acc) -> if(f.(h), do: [h|acc], else: acc) end)

  @type acc :: any
  @spec reduce(list, acc, ((any, acc) -> acc)) :: acc
  def reduce(l, acc, f) do
    case l do
      []    -> acc
      [h|t] -> reduce(t,f.(h,acc),f)
    end
  end

  @spec append(list, list) :: list
  def append(a, b) do
    case a do
      []    -> b
      [h|t] -> [h|append(t,b)]
    end
  end

  @spec concat([[any]]) :: [any]
  def concat(ll), do: _concat([],ll) |> reverse
  defp _concat(result,ll) do
    case ll do
      []    -> result
      [h|t] -> _concat(result,h) |> _concat t
      h     -> [h|result]
    end
  end
end
