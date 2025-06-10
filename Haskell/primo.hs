-- Exercício 2) Dado um inteiro positivo n, verificar se n é primo.

main :: IO ()
main = do
    putStrLn "Digite um número inteiro positivo n:"
    nStr <- getLine
    let n = read nStr :: Integer

    if n <= 0 then
        putStrLn "Erro: n deve ser um número inteiro positivo."
    else if isPrime n then
        putStrLn (show n ++ " é primo.")
    else
        putStrLn (show n ++ " não é primo.")

isPrime :: Integer -> Bool
isPrime 1 = False
isPrime 2 = True
isPrime n = not (any (\x -> n `mod` x == 0) [2..(floor . sqrt . fromIntegral) n])
