function vec3Add(x, y) {
    var added = [];
    for(var i = 0; i < x.length; i++) {
        added.push(x[i] + y[i]);
    }
    return added;
}

function vec3Sub(x, y) {
    var subbed = [];
    for(var i = 0; i < x.length; i++) {
        subbed.push(x[i] - y[i]);
    }
    return subbed;
} 

function vec3Scale(x, y) {
    var scaled = [];
    for(var i = 0; i < x.length; i++) {
        scaled.push(x[i] * y);
    }
    return scaled;
}  

function vec3Mag(v) {
    var magnitude = 0;
    for(var i = 0; i < v.length; i++) {
        magnitude += Math.pow(v[i], 2);
    }
    return Math.sqrt(magnitude);
}

function dotProduct(x, y) {
    var dotted = [];
    for(var i = 0; i < x.length; i++) {
        dotted.push(x[i] * y[i]);
    }
    return dotted;
}

function crossProduct(x, y) {
    var cx =  (x[1] * y[2]) - (x[2] * y[1]);
    var cy =  (x[2] * y[0]) - (x[0] * y[2]);
    var cz =  (x[0] * y[1]) - (x[1] * y[0]);
    return [cx, cy, cz];
}