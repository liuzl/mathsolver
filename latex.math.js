var nf = nf || {};
var vf = vf || {};
nf.math = {};
nf.util = {};
vf.be = {};

nf.list = function(type, cnt) {
    //TODO
}

nf.it = function() {}
nf.what = function() {}

nf.math.expression = function(s) {
    return s.split("").join('*');
}

nf.math.to_number = function(s) {
    return Number(s);
}

nf.math.decimal = function(s) {
    s = s.toString();
    var n = Number(s);
    return n / Math.pow(10, s.length);
}

nf.math.sum = function(x, y) {
    //return x + y;
    return "("+x+"+"+y+")";
}

nf.math.sub = function(x, y) {
    //return x - y;
    return "("+x+"-"+y+")";
}

nf.math.mul = function(x, y) {
    //return x * y;
    return "("+x+"*"+y+")";
}

nf.math.div = function(x, y) {
    //return x / y;
    return "("+x+"/"+y+")";
}

nf.math.neg = function(x) {
    return -x;
}

nf.math.pow = function(x, y) {
    //return Math.pow(x, y);
    return "("+x+"^"+y+")";
}

nf.util.concat = function(x, y) {
    return x.toString() + y.toString();
}

vf.be.eq = function(x, y) {
    return x + "=" + y;
}

vf.be.neq = function(x, y) {
    return x + "!=" + y;
}

vf.be.gt = function(x, y) {
    return x + ">" + y;
}

vf.be.geq = function(x, y) {
    return x + ">=" + y;
}

vf.be.lt = function(x, y) {
    return x + "<" + y;
}

vf.be.leq = function(x, y) {
    return x + "<=" + y;
}

nf.math.func = function(name, x, y) {
    return name+"("+x+","+y+")";
}
