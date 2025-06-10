-- Exercício 1) Dados um inteiro x e um inteiro não-negativo n, calcular 𝑥^𝑛.


main :: IO ()
main = do
    putStrLn "Digite um número inteiro x:"
    xStr <- getLine
    putStrLn "Digite um número inteiro não-negativo n:"
    nStr <- getLine

    let x = read xStr :: Integer
    let n = read nStr :: Integer

    if n < 0 then
        putStrLn "Erro: n deve ser um número não-negativo."
    else
        putStrLn ("Resultado: " ++ show (power x n))

power :: Integer -> Integer -> Integer
power _ 0 = 1
power x n = x * power x (n - 1)