import Data.Monoid

add :: Num a => a -> a -> a
add x y = getSum $ Sum x <> Sum y

main :: IO ()
main = do
    let result = add 5 7
    putStrLn $ "Result: " ++ show result
