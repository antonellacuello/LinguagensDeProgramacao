-- Exercício 4) Escreva um programa que converta temperaturas de Fahrenheit para Celsius, seguindo a equação 𝐶=(𝐹−32)×5/9.

main :: IO ()
main = do
    putStrLn "Digite a temperatura em Fahrenheit:"
    input <- getLine
    let fahrenheit = read input :: Float
    let celsius = (fahrenheit - 32) * 5 / 9
    putStrLn ("A temperatura em Celsius é: " ++ show celsius)