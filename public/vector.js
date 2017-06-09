function vec3Add(a, b) {
    return {x: (a.x + b.x), y: (a.y + b.y), z: (a.z + b.z)};
}

function vec3Sub(a, b) {
    return {x: (a.x - b.x), y: (a.y - b.y), z: (a.z - b.z)};
} 

function vec3Scale(a, f) {
    return {x: (a.x * f), y: (a.y * f), z: (a.z * f)};
}  

function vec3Mag(a) {
    var mag = Math.pow(a.x, 2) + Math.pow(a.y, 2) + Math.pow(a.z, 2);
    return Math.sqrt(mag);
}

function vec3Dot(a, b) {
    return (a.x * b.x) + (a.y * b.y) + (a.z * b.z);
}

function vec3Cross(a, b) {
    var cx =  (a.y * b.z) - (a.z * b.y);
    var cy =  (a.z * b.x) - (a.x * b.z);
    var cz =  (a.x * b.y) - (a.y * b.x);
    return {x: cx, y: cy, z: cz};
}

function vec3Yaw(a) {
    return Math.atan2(a.z, a.x);
}

function vec3Pitch(a) {
    return Math.asin(a.y);
}