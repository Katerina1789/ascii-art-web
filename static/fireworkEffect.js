/*
This javascript file handles the firework effect interactivity and grid generation for desplaying the colors

The entire backround is separated into a 2 dimensional grid of square blocks of "pixels" a grid of (40 x 40) and each block is 40 pixels wide and 40 pixels tall.

This give the ability to calculate the distance and add effects to every individual square block
although we just use it to calcute what box will be colored
*/


//you can learn more about Canvas in the following links 
//Detailed-Guide:https://joshondesign.com/p/books/canvasdeepdive/chapter01.html#paths
//Cheat-Sheet:https://websitesetup.org/wp-content/uploads/2015/11/Infopgraphic-CanvasCheatSheet-Final2.pdf
//MDN:https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/canvas
const gridCanvas = document.getElementById('gridCanvas');
const gridCtx = gridCanvas.getContext('2d');

//Get the dimensions of the window and set the canvas to match it, 
//this will make sure the canvas covers the entire background (We dont not optimize for resize of the window)
gridCanvas.width = window.innerWidth;
gridCanvas.height = window.innerHeight;

// Make sure canvases are visible (though CSS should handle this, JS redundancy doesn't hurt if we removed the toggler)
document.getElementById('gridCanvas').style.display = 'block';
document.getElementById('fireworksCanvas').style.display = 'block';


// Box size used for generated the grid of the boxes and to define the size of each individual box
const BOX_SIZE = 40;
//We are stroring the boxes inside boxes and we use the keyword "let" to specifically trap the usability of this object inside this 
//file in case we need to use it differently in another part of our programme in this case ("File-Scope").
let boxes = {};


//Initiating the Grid taking the height and lenght of the window 
function initGrid() {
    gridCanvas.width = window.innerWidth;
    gridCanvas.height = window.innerHeight;
    boxes = {};
    drawGrid();
}


//Draws the grid by filling with black collor and outlines/seperates the blocks with grey
function drawGrid() {
    gridCtx.fillStyle = '#000';
    gridCtx.fillRect(0, 0, gridCanvas.width, gridCanvas.height);

    gridCtx.strokeStyle = '#333';
    gridCtx.lineWidth = 1;
    //Getting the grid width in pixels and we iderate box by box by adding every time the BOX_SIZE=40 "40px"
    //We reapeting this for the height
    for (let x = 0; x < gridCanvas.width; x += BOX_SIZE) {
        for (let y = 0; y < gridCanvas.height; y += BOX_SIZE) {
            //Creates a 4 point connection to the grid as a baseline to color it later  
            gridCtx.strokeRect(x, y, BOX_SIZE, BOX_SIZE);

            /*Creates a "Key"  to map every tile 40x40 with an id 
            For ease of identification in case of a check or debug 
            Ex: Lets say you have trable with the box in starting position x/y:120 
            These numbers don't give you a clear representation of the position that the box is locaded
            Unless you manualy divide this number with 40 (The max  size of each cordinate)
            */
            const key = `${Math.floor(x / BOX_SIZE)},${Math.floor(y / BOX_SIZE)}`;
            if (boxes[key]) {
                gridCtx.fillStyle = boxes[key];
                //We need to fill the boxes pixel by pixel fillRect does that.
                //https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/fillRect
                gridCtx.fillRect(x, y, BOX_SIZE, BOX_SIZE);
            }
        }
    }
}
//This function is used to fill the boxes with the selected color
function fillBox(x, y, color) {
    const boxX = Math.floor(x / BOX_SIZE);
    const boxY = Math.floor(y / BOX_SIZE);
    const key = `${boxX},${boxY}`;
    boxes[key] = color;

    // Direct render: Only draw the one box that changed
    gridCtx.fillStyle = color;
    gridCtx.fillRect(boxX * BOX_SIZE, boxY * BOX_SIZE, BOX_SIZE, BOX_SIZE);

    // Redraw the border for just this box so it matches the grid
    gridCtx.strokeStyle = '#333';
    gridCtx.strokeRect(boxX * BOX_SIZE, boxY * BOX_SIZE, BOX_SIZE, BOX_SIZE);
}


//Listener for resizing the window to make sure the grid is always covering the entire background and to reinitiate the grid with the new dimensions
window.addEventListener('resize', initGrid);
initGrid();

function animateGrid() {

    requestAnimationFrame(animateGrid);
}

animateGrid();



// FIREWORKS EFFECT

/*We creating another layer on top of our first layer to see out fireworks in the backround */
const fireworksCanvas = document.getElementById('fireworksCanvas');
const fireworksCtx = fireworksCanvas.getContext('2d');


//Taking the dimensions of the window and set the canvas to match it
fireworksCanvas.width = window.innerWidth;
fireworksCanvas.height = window.innerHeight;

let particles = [];
let rockets = [];

// Rocket class represents the projectile that travels from the bottom of the screen to the target point where it will explode
class Rocket {
    //The constructor takes the target x and y coordinates where the firework will explode, and initializes the rocket's position, color, speed, and wobble.
    constructor(targetX, targetY) {
        this.x = targetX;
        this.y = fireworksCanvas.height;
        this.targetY = targetY;
        this.color = `hsl(${Math.random() * 360}, 100%, 50%)`;
        this.speed = Math.random() * 1 + 8;
        this.wobble = Math.random() * 2 - 1;
    }

    // Draws the rocket as a small circle on the fireworks canvas
    draw() {
        fireworksCtx.beginPath();
        fireworksCtx.fillStyle = this.color;
        fireworksCtx.arc(this.x, this.y, 3, 0, Math.PI * 2);
        fireworksCtx.fill();
    }

    // Handles the rocket's movement, increasing its height (decreasing Y) and adding a slight wobble effect
    // It also checks if the rocket has reached its destination to trigger the explosion
    update(index) {
        this.y -= this.speed;
        this.x += Math.sin(this.y * 0.1) * 0.5;
        //When the rocket meets the destination
        if (this.y <= this.targetY) {
            rockets.splice(index, 1);
            fillBox(this.x, this.targetY, this.color);
            createExplosion(this.x, this.targetY, this.color);
        }
    }
}

// Particle class represents the individual sparks that fly out during a firework explosion
class Particle {
    // Initializes the particle with a position, color, and random velocity to create a radial explosion effect
    constructor(x, y, color) {
        this.x = x;
        this.y = y;
        this.color = color;

        const angle = Math.random() * Math.PI * 2;
        const velocity = Math.random() * 5 + 2;

        this.vx = Math.cos(angle) * velocity;
        this.vy = Math.sin(angle) * velocity;

        this.alpha = 1;
        this.friction = 0.98;
        this.gravity = 0.05;
    }

    // Renders the particle on the canvas, using the alpha property to create a fading out visual
    draw() {
        fireworksCtx.save();
        // Set the global alpha to create a fading effect as the particle moves
        fireworksCtx.globalAlpha = this.alpha;
        //We start drawing 
        fireworksCtx.beginPath();

        fireworksCtx.arc(this.x, this.y, 2, 0, Math.PI * 2);
        fireworksCtx.fillStyle = this.color;
        fireworksCtx.fill();
        fireworksCtx.restore();
    }

    // Updates the particle's physics: applies friction to slow it down, gravity to pull it down, 
    // and decreases its transparency until it's completely faded out
    update(index) {
        this.vx *= this.friction;
        this.vy *= this.friction;
        this.vy += this.gravity;
        this.x += this.vx;
        this.y += this.vy;
        this.alpha -= 0.015;

        if (this.alpha <= 0) {
            particles.splice(index, 1);
        }
    }
}
//Creates the explotion with the particles taking as input the position and color
function createExplosion(x, y, color) {
    const particleCount = 50;
    for (let i = 0; i < particleCount; i++) {
        particles.push(new Particle(x, y, color));
    }
}
//
function animateFireworks() {
    requestAnimationFrame(animateFireworks);
    // Clear the fireworks canvas with a slight opacity to create a fading trail effect
    fireworksCtx.globalCompositeOperation = 'destination-out';
    fireworksCtx.fillStyle = 'rgba(0, 0, 0, 0.3)';
    //This will create a fading effect for the fireworks, as it will slowly clear the 
    //previous frames while still allowing the new particles to be visible. 
    //The 'destination-out' composite operation ensures that we are only clearing the existing pixels without affecting the new drawings.
    fireworksCtx.fillRect(0, 0, fireworksCanvas.width, fireworksCanvas.height);
    fireworksCtx.globalCompositeOperation = 'source-over';
    //Update and draw rockets and particles, 
    //iterating backwards to safely remove elements from the arrays while iterating
    for (let i = rockets.length - 1; i >= 0; i--) {
        rockets[i].draw();
        rockets[i].update(i);
    }

    for (let i = particles.length - 1; i >= 0; i--) {
        particles[i].draw();
        particles[i].update(i);
    }
}

animateFireworks();


window.addEventListener('click', (event) => {
    rockets.push(new Rocket(event.clientX, event.clientY));
});


//Sets the time for the firework to be generated every 1 second and it will be generated
setInterval(() => {
    let x;
    // Randomly decide to generate the firework on the left or right side of the screen
    // This creates a more dynamic and visually interesting effect, 
    // as fireworks will appear from both sides of the screen rather than just one.
    if (Math.random() < 0.5) {
        x = Math.random() * (fireworksCanvas.width * 0.3);
    } else {
        x = (fireworksCanvas.width * 0.7) + Math.random() * (fireworksCanvas.width * 0.3);
    }
    // Randomly generate the height of the firework between 10% and 70% of the screen height
    const y = Math.random() * (fireworksCanvas.height * 0.6) + (fireworksCanvas.height * 0.1);
    rockets.push(new Rocket(x, y));
}, 1000);

//Listener for resizing the window to make sure the fireworks canvas is 
// always covering the entire background and to reinitiate the fireworks 
// canvas with the new dimensions.
window.addEventListener('resize', () => {
    fireworksCanvas.width = window.innerWidth;
    fireworksCanvas.height = window.innerHeight;
});


// Copy to Clipboard Functionality
const copyBtn = document.getElementById('copyBtn');
const asciiResult = document.getElementById('asciiResult');

if (copyBtn && asciiResult) {
    copyBtn.addEventListener('click', () => {
        const textToCopy = asciiResult.innerText;
        navigator.clipboard.writeText(textToCopy).then(() => {
            const originalText = copyBtn.textContent;
            copyBtn.textContent = 'Copied!';
            copyBtn.style.background = 'rgba(255, 255, 255, 0.3)';

            setTimeout(() => {
                copyBtn.textContent = originalText;
                copyBtn.style.background = '';
            }, 1000);
        }).catch(err => {
            console.error('Failed to copy text: ', err);
            copyBtn.textContent = 'Error';
        });
    });
}
