window.onload = function() {

    function createSphere() {
        var geometry = new THREE.SphereGeometry(300, 9, 9);
        var material = new THREE.MeshLambertMaterial({color: 0x6313E5});
        return new THREE.Mesh(geometry, material);
    }

    function createBox() {
        var geometry = new THREE.BoxGeometry(500, 500, 1000);
        var material = new THREE.MeshLambertMaterial( {color: 0x00ff00} );
        return new THREE.Mesh(geometry, material);
    }

    var meshes = {};

    var renderer = new THREE.WebGLRenderer({canvas: document.getElementById("screen"), antialias: true});
    renderer.setClearColor(0x260C37);
    renderer.setPixelRatio(window.devicePixelRatio);
    renderer.setSize(window.innerWidth, window.innerHeight);

    var camera = new THREE.PerspectiveCamera(35, window.innerWidth / window.innerHeight, 0.1, 450000);
    //camera.up = new THREE.Vector3(0, 1, 0);
    //camera.lookAt(new THREE.Vector3(0, 0, 1));
    camera.position.set(0, 0, 3000);

    var vecty = new THREE.Vector3(1, 0, -1);
    vecty.applyQuaternion(camera.quaternion);
    camera.lookAt(vecty); 

    var scene = new THREE.Scene();

    var light_ambient = new THREE.AmbientLight(0xffffff, 0.8);
    scene.add(light_ambient);

    var light_point = new THREE.PointLight(0xffffff, 0.9);
    scene.add(light_point);

    var particle_material = new THREE.MeshLambertMaterial();
    particle_material.map = THREE.TextureLoader("tex_space_ship.jpg");
    particle_material.side = THREE.DoubleSide;

    var loader = new THREE.JSONLoader();

    function createArwing(id) {
        loader.load("arwing.json", 
        function(geometry) {
            var texture = THREE.ImageUtils.loadTexture("tex_space_ship.jpg");
            var material = new THREE.MeshBasicMaterial({map: texture});
            var mesh = new THREE.Mesh(geometry, material);
            mesh.scale.set(110, 110, 110);
            
            meshes[id] = mesh;
            scene.add(mesh);}
        );
    }

    var pressed_keys = []
    for(var i = 0; i < 222; i++) {
        pressed_keys.push(false);
    }

    function keyAction() {
        var camera_speed = 45;

        var z_vector = Math.cos(camera.rotation.y) * Math.cos(camera.rotation.x);
        var y_vector = Math.sin(camera.rotation.x);
        var x_vector = Math.sin(camera.rotation.y) * Math.cos(camera.rotation.x);

        var forward = {x: x_vector, y: y_vector, z: z_vector};
        var up = [0, 1, 0];
        var right = vec3Cross(forward, up);
        up = vec3Cross(forward, right);

        forward = vec3Scale(forward, camera_speed);
        right = vec3Scale(right, camera_speed);
        up = vec3Scale(up, camera_speed);
    }

    function pressKey(evt) {
       var key = parseInt(evt.keyCode);
       pressed_keys[key] = true;
    }      

    function releaseKey(evt) {
        var key = parseInt(evt.keyCode);
        pressed_keys[key] = false;
    }

    var last_mouse_x = 100;
    var last_mouse_y = 100;
    var first_mouse = true;

    function mouseMotion(evt) {
        if(first_mouse) {
            last_mouse_x = evt.clientX;
            last_mouse_y = evt.clientY;
            first_mouse = false;
        }
        var x = evt.clientX;
        var y = evt.clientY;

        var x_offset = x - last_mouse_x;
        var y_offset = y - last_mouse_y;

        var sensitivity = 0.003;

        //camera.rotation.y -= x_offset * sensitivity;
        //camera.rotation.x -= y_offset * sensitivity;

        last_mouse_x = x;
        last_mouse_y = y;
    }

    function onWindowResize() {
        camera.aspect = window.innerWidth / window.innerHeight;
        camera.updateProjectionMatrix();
        renderer.setSize( window.innerWidth, window.innerHeight );
    }

    document.onkeydown = pressKey;
    document.onkeyup = releaseKey;
    document.onmousemove = mouseMotion;
    window.addEventListener("resize", onWindowResize, false);

    var socket = new WebSocket("ws://localhost:8000/ws");
    socket.binaryType = "arraybuffer";

    var my_id = -1;
    var world_data = [];

    function updateWorld(ents) {
        for (var i = 0; i < ents.length; i++) {
            var ent = JSON.parse(ents[i]);
            var id = ent.id;

            if (!(id in meshes)) {
                var sphere = createBox();
                meshes[id] = sphere;
                scene.add(sphere);
            }

            var mesh = meshes[id];
            mesh.position.set(ent.pos.x, ent.pos.y, ent.pos.z);
            mesh.rotation.set(ent.euler.x, ent.euler.y, ent.euler.z, "YXZ");

            if (ent.type === "player") {
                camera.position.set(0,0,0);
                //camera.position.set(mesh.position.x - (ent.forward.x * 4500), mesh.position.y - (ent.forward.y * 4500) + 1000, mesh.position.z - (ent.forward.z * 4500));
                //camera.rotation.set(ent.euler.x, ent.euler.y, ent.euler.z, "YXZ");
                camera.lookAt(new THREE.Vector3(ent.forward.x, ent.forward.y, ent.forward.z));
                camera.up.set(ent.up.x, ent.up.y, ent.up.z);
                camera.position.set(mesh.position.x, mesh.position.y, mesh.position.z);
            }

            if (ent.type == "animal") {
                mesh.material.color.setHex(0x6313E5);
            } else {
                if (ent.type == "plant") {
                    mesh.material.color.setHex(0x3EC70E);
                } else {
                    mesh.material.color.setHex(0xC70E0E);
                }
            }
        }
    }

    socket.onopen = function() {
        console.log("connected to server");
    }

    socket.onclose = function(e) {
        console.log("connection closed (" + e.code + ")");
    }

    socket.onmessage = function(message) {
        var msg = JSON.parse(message.data);
        if (msg.type == "id") {
            my_id = msg.data;
        }
        if (msg.type == "update") {
            world_data = msg.data;
        }
    }

    function sendPressed() {
        var msg = 0;
        if(pressed_keys[87]) { msg = setBitAt(msg, 0); } // W
        if(pressed_keys[83]) { msg = setBitAt(msg, 1); } // S
        if(pressed_keys[65]) { msg = setBitAt(msg, 2); } // A
        if(pressed_keys[68]) { msg = setBitAt(msg, 3); } // D
        if(pressed_keys[81]) { msg = setBitAt(msg, 4); } // Q
        if(pressed_keys[69]) { msg = setBitAt(msg, 5); } // E

        if(pressed_keys[73]) { msg = setBitAt(msg, 6); } // I
        if(pressed_keys[75]) { msg = setBitAt(msg, 7); } // K
        if(pressed_keys[74]) { msg = setBitAt(msg, 8); } // J
        if(pressed_keys[76]) { msg = setBitAt(msg, 9); } // L

        if(pressed_keys[85]) { msg = setBitAt(msg, 10); } // U
        if(pressed_keys[79]) { msg = setBitAt(msg, 11); } // O
        
        msg_arr = [msg];
        var data = new Uint8Array(2);
        data[0] = msg & 255;
        data[1] = msg >> 8;
        socket.send(data.buffer);
    }

    function setBitAt(val, pos) {
        var mask = 1 << pos;
        return val | mask;
    }

    function render() {
        keyAction();
        if (socket.readyState == 1) {
            sendPressed();
        }
        updateWorld(world_data);
        renderer.render(scene, camera);
        requestAnimationFrame(render);
    }

    requestAnimationFrame(render);
}
















