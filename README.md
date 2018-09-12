# mathsolver

## 示例

* NL
```
五与2.8的和的平方的立方的1.5次方与two的和
````

* FMR
```js
nf.math.sum(
  nf.math.pow(
    nf.math.pow(
      nf.math.pow(
        nf.math.sum(
          5,
          nf.math.to_number(
            "2.8"
          )
        ),
        2
      ),
      3
    ),
    nf.math.to_number("1.5")
  ),
  2
);
```

* Denotation
```
106868922.9132846
```