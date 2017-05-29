
window.onload = function() 
{
    var pressed_keys = []
    for(var i = 0; i < 222; i++)
    {
        pressed_keys.push(false);
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

    //var sphere = createSphere();
    //scene.add(sphere);

    function createSphere()
    {
        var geometry = new THREE.SphereGeometry(300, 9, 9);
        var material = new THREE.MeshLambertMaterial({color: 0x6313E5});
        return new THREE.Mesh(geometry, material);
    }

    var particle_material = new THREE.MeshLambertMaterial();
    particle_material.map = THREE.TextureLoader("tex_space_ship.jpg");
    particle_material.side = THREE.DoubleSide;

    var loader = new THREE.JSONLoader();

    function createArwing(id)
    {
        loader.load("arwing.json", 
        function(geometry) 
        {
            var texture = THREE.ImageUtils.loadTexture("tex_space_ship.jpg");
            var material = new THREE.MeshBasicMaterial({map: texture});
            //var material = new THREE.MultiMaterial(materials);
            var mesh = new THREE.Mesh(geometry, material);
            mesh.scale.set(65, 65, 65);
            
            meshes[id] = mesh;
            scene.add(mesh);
        }
        );
    }


    function render() 
    {
        keyAction();

        //createArwings();

        renderer.render(scene, camera);
        requestAnimationFrame(render);
    }

    var last_mouse_x = 100;
    var last_mouse_y = 100;

    var first_mouse = true;

    function mouseMotion(evt)
    {
        if(first_mouse)
        {
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

    function keyAction()
    {
        var camera_speed = 45;

        var z_vector = Math.cos(camera.rotation.y) * Math.cos(camera.rotation.x);
        var y_vector = Math.sin(camera.rotation.x);
        var x_vector = Math.sin(camera.rotation.y) * Math.cos(camera.rotation.x);

        var forward = [x_vector, y_vector, z_vector];
        var up = [0, 1, 0];
        var right = crossProduct(forward, up);
        up = crossProduct(forward, right);

        forward = vec3Scale(forward, camera_speed);
        right = vec3Scale(right, camera_speed);
        up = vec3Scale(up, camera_speed);
     
        if(pressed_keys[87]) // W
        {
            camera.position.x -= forward[0];
            camera.position.y += forward[1];
            camera.position.z -= forward[2];
        }
        if(pressed_keys[83]) // S
        {
            camera.position.x += forward[0];
            camera.position.y -= forward[1];
            camera.position.z += forward[2];
        }
        if(pressed_keys[65]) // A
        {
            camera.position.x += right[0];
            camera.position.y += right[1];
            camera.position.z += right[2];
        }
        if(pressed_keys[68]) // D
        {
            camera.position.x -= right[0];
            camera.position.y -= right[1];
            camera.position.z -= right[2];
        }
        if(pressed_keys[81]) // Q
        {
            camera.position.x -= up[0];
            camera.position.y -= up[1];
            camera.position.z += up[2];
        }
        if(pressed_keys[69]) // E
        {
            camera.position.x += up[0];
            camera.position.y += up[1];
            camera.position.z -= up[2];
        }

        if(pressed_keys[38]) // up
        {
            camera.rotation.x += 0.01
        }
        if(pressed_keys[40]) // down
        {
            camera.rotation.x -= 0.01;
        }
        if(pressed_keys[37]) // left
        {
            camera.rotation.y += 0.01;
        }
        if(pressed_keys[39]) // right
        {
            camera.rotation.y -= 0.01;
        }
    }

    function pressKey(evt)
    {
       var key = parseInt(evt.keyCode);
       pressed_keys[key] = true;
    }      

    function releaseKey(evt)
    {
        var key = parseInt(evt.keyCode);
        pressed_keys[key] = false;
    }

    function onWindowResize() 
    {
        camera.aspect = window.innerWidth / window.innerHeight;
        camera.updateProjectionMatrix();
        renderer.setSize( window.innerWidth, window.innerHeight );
    }

    function vec3Add(x, y)
    {
        var added = [];
        for(var i = 0; i < x.length; i++)
        {
            added.push(x[i] + y[i]);
        }
        return added;
    }

    function vec3Sub(x, y)
    {
        var subbed = [];
        for(var i = 0; i < x.length; i++)
        {
            subbed.push(x[i] - y[i]);
        }
        return subbed;
    } 

    function vec3Scale(x, y)
    {
        var scaled = [];
        for(var i = 0; i < x.length; i++)
        {
            scaled.push(x[i] * y);
        }
        return scaled;
    }  

    function vec3Mag(v)
    {
        var magnitude = 0;
        for(var i = 0; i < v.length; i++)
        {
            magnitude += Math.pow(v[i], 2);
        }
        return Math.sqrt(magnitude);
    }

    function dotProduct(x, y)
    {
        var dotted = [];
        for(var i = 0; i < x.length; i++)
        {
            dotted.push(x[i] * y[i]);
        }

        return dotted;
    }

    function crossProduct(x, y)
    {
        var cx =  (x[1] * y[2]) - (x[2] * y[1]);
        var cy =  (x[2] * y[0]) - (x[0] * y[2]);
        var cz =  (x[0] * y[1]) - (x[1] * y[0]);

        return [cx, cy, cz];
    }

    requestAnimationFrame(render);



    document.onkeydown = pressKey;
    document.onkeyup = releaseKey;
    document.onmousemove = mouseMotion;

    window.addEventListener("resize", onWindowResize, false);

    socket = new WebSocket("ws://localhost:8000/ws");

    socket.onopen = function()
    {
        console.log("connected to server");
    }

    socket.onclose = function(e)
    {
        console.log("connection closed (" + e.code + ")");
    }

    socket.onmessage = function(message)
    {
        var animals = JSON.parse(message.data);
        for(var i = 0; i < animals.length; i++)
        {
            var ani = JSON.parse(animals[i]);
            //console.log(ani);
            var id = ani.id;

            if(!(id in meshes))
            {
                var sphere = createSphere();
                meshes[id] = sphere;
                scene.add(sphere);
                //createArwing(id);
            }

            var mesh = meshes[id];
            mesh.position.x = ani.pos.x;
            mesh.position.y = ani.pos.y;
            mesh.position.z = ani.pos.z;

            //sphere.position.x = ani.target.x;
            //sphere.position.y = ani.target.y;
            //sphere.position.z = ani.target.z;

            //mesh.rotation.y = ani.euler.yaw
        }
    }

    function send()
    {
        var message = document.getElementById("message").value;
        socket.send(message);
    }
}
















