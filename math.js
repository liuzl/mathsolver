var nf = nf || {};
nf.math = {};

nf.math.to_number = function(s) {
    return Number(s);
}

nf.math.sum = function(x, y) {
    return x + y;
}

nf.math.sub = function(x, y) {
    return x - y;
}

nf.math.mul = function(x, y) {
    return x * y;
}

nf.math.div = function(x, y) {
    return x / y;
}

nf.math.neg = function(x) {
    return -x;
}

nf.math.pow = function(x, y) {
    return Math.pow(x, y);
}

