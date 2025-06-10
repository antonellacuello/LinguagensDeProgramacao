-- Exercício 3) Dado um número inteiro positivo n, imprimir os n primeiros naturais ímpares.

main :: IO ()
main = do
    putStrLn "Digite um número inteiro positivo:"
    input <- getLine
    let n = read input :: Int
    if n > 0
        then printOddUpToN n
        else putStrLn "Por favor, digite um número inteiro positivo."

printOddUpToN :: Int -> IO ()
printOddUpToN n = do
    let oddNumbers = [x | x <- [1..n], odd x]
    putStrLn "Números naturais ímpares até n:"
    putStrLn $ formatList oddNumbers

formatList :: [Int] -> String
formatList = init . unwords . map (++",") . map show