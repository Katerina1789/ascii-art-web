let activeEffect = 'warp';

document.querySelectorAll('input[name="effect"]').forEach(radio => {
    radio.addEventListener('change', (e) => {
        activeEffect = e.target.value;
        updateEffects();
    });
});

function updateEffects() {
    const warpCanvas = document.getElementById('warpCanvas');
    const gridCanvas = document.getElementById('gridCanvas');
    const fireworksCanvas = document.getElementById('fireworksCanvas');

    if (activeEffect === 'warp') {
        warpCanvas.style.display = 'block';
        gridCanvas.style.display = 'none';
        fireworksCanvas.style.display = 'none';
    } else if (activeEffect === 'fireworks') {
        warpCanvas.style.display = 'none';
        gridCanvas.style.display = 'block';
        fireworksCanvas.style.display = 'block';
    }
}

// GRID BACKGROUND
const gridCanvas = document.getElementById('gridCanvas');
const gridCtx = gridCanvas.getContext('2d');
gridCanvas.width = window.innerWidth;
gridCanvas.height = window.innerHeight;

const BOX_SIZE = 40;
let boxes = {};

function initGrid() {
    gridCanvas.width = window.innerWidth;
    gridCanvas.height = window.innerHeight;
    boxes = {};
    drawGrid();
}

function drawGrid() {
    gridCtx.fillStyle = '#000';
    gridCtx.fillRect(0, 0, gridCanvas.width, gridCanvas.height);

    gridCtx.strokeStyle = '#333';
    gridCtx.lineWidth = 1;

    for (let x = 0; x < gridCanvas.width; x += BOX_SIZE) {
        for (let y = 0; y < gridCanvas.height; y += BOX_SIZE) {
            gridCtx.strokeRect(x, y, BOX_SIZE, BOX_SIZE);

            const key = `${Math.floor(x / BOX_SIZE)},${Math.floor(y / BOX_SIZE)}`;
            if (boxes[key]) {
                gridCtx.fillStyle = boxes[key];
                gridCtx.fillRect(x, y, BOX_SIZE, BOX_SIZE);
            }
        }
    }
}

function fillBox(x, y, color) {
    const boxX = Math.floor(x / BOX_SIZE);
    const boxY = Math.floor(y / BOX_SIZE);
    const key = `${boxX},${boxY}`;
    boxes[key] = color;
    drawGrid();
}

window.addEventListener('resize', initGrid);
initGrid();

function animateGrid() {
    drawGrid();
    requestAnimationFrame(animateGrid);
}

animateGrid();

// WARP EFFECT
const warpCanvas = document.getElementById('warpCanvas');
const warpCtx = warpCanvas.getContext('2d');

let warpWidth, warpHeight, warpCenterX, warpCenterY;

function resizeWarp() {
    warpWidth = warpCanvas.width = window.innerWidth;
    warpHeight = warpCanvas.height = window.innerHeight;
    warpCenterX = warpWidth / 2;
    warpCenterY = warpHeight / 2;
}

window.addEventListener('resize', resizeWarp);
resizeWarp();

const STAR_COUNT = 400;
const SPEED = 0.1;
const STAR_SIZE = 2;
const SPREAD = 800;

let stars = [];

class Star {
    constructor() {
        this.reset();
    }

    reset() {
        this.x = (Math.random() - 0.5) * SPREAD * 2;
        this.y = (Math.random() - 0.5) * SPREAD * 2;
        this.z = (Math.random() * SPREAD);
        this.pz = this.z;
    }

    update() {
        this.z = this.z - (SPEED * 20);
        if (this.z <= 0) {
            this.reset();
            this.z = SPREAD;
            this.pz = this.z;
        }
    }

    draw() {
        const sx = (this.x / this.z) * 100 + warpCenterX;
        const sy = (this.y / this.z) * 100 + warpCenterY;
        const size = (1 - this.z / SPREAD) * STAR_SIZE * 3;

        if (this.z < 1) return;

        warpCtx.beginPath();
        const opacity = (1 - this.z / SPREAD);
        warpCtx.fillStyle = `rgba(255, 255, 255, ${opacity})`;
        warpCtx.arc(sx, sy, size, 0, Math.PI * 2);
        warpCtx.fill();
    }
}

for (let i = 0; i < STAR_COUNT; i++) {
    stars.push(new Star());
}

function animateWarp() {
    warpCtx.fillStyle = "black";
    warpCtx.fillRect(0, 0, warpWidth, warpHeight);

    for (let star of stars) {
        star.update();
        star.draw();
    }

    requestAnimationFrame(animateWarp);
}

animateWarp();

// FIREWORKS EFFECT
const fireworksCanvas = document.getElementById('fireworksCanvas');
const fireworksCtx = fireworksCanvas.getContext('2d');

fireworksCanvas.width = window.innerWidth;
fireworksCanvas.height = window.innerHeight;

let particles = [];
let rockets = [];

class Rocket {
    constructor(targetX, targetY) {
        this.x = targetX;
        this.y = fireworksCanvas.height;
        this.targetY = targetY;
        this.color = `hsl(${Math.random() * 360}, 100%, 50%)`;
        this.speed = 12;
        this.wobble = Math.random() * 2 - 1;
    }

    draw() {
        fireworksCtx.beginPath();
        fireworksCtx.fillStyle = this.color;
        fireworksCtx.arc(this.x, this.y, 3, 0, Math.PI * 2);
        fireworksCtx.fill();
    }

    update(index) {
        this.y -= this.speed;
        this.x += Math.sin(this.y * 0.1) * 0.5;

        if (this.y <= this.targetY) {
            rockets.splice(index, 1);
            fillBox(this.x, this.targetY, this.color);
            createExplosion(this.x, this.targetY, this.color);
        }
    }
}

class Particle {
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

    draw() {
        fireworksCtx.save();
        fireworksCtx.globalAlpha = this.alpha;
        fireworksCtx.beginPath();
        fireworksCtx.arc(this.x, this.y, 2, 0, Math.PI * 2);
        fireworksCtx.fillStyle = this.color;
        fireworksCtx.fill();
        fireworksCtx.restore();
    }

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

function createExplosion(x, y, color) {
    const particleCount = 100;
    for (let i = 0; i < particleCount; i++) {
        particles.push(new Particle(x, y, color));
    }
}

function animateFireworks() {
    requestAnimationFrame(animateFireworks);

    fireworksCtx.globalCompositeOperation = 'destination-out';
    fireworksCtx.fillStyle = 'rgba(0, 0, 0, 0.2)';
    fireworksCtx.fillRect(0, 0, fireworksCanvas.width, fireworksCanvas.height);
    fireworksCtx.globalCompositeOperation = 'source-over';

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

let clickCount = 0;
const fillAllBtn = document.getElementById('fillAllBtn');

window.addEventListener('click', (event) => {
    if (activeEffect === 'fireworks') {
        // Only count clicks on the canvas/background, not the button itself
        if (event.target !== fillAllBtn) {
            clickCount++;
            rockets.push(new Rocket(event.clientX, event.clientY));

            if (clickCount >= 20) {
                fillAllBtn.style.display = 'block';
            }
        }
    }
});

fillAllBtn.addEventListener('click', (e) => {
    e.stopPropagation();

    if (fillAllBtn.textContent === 'Fill All Boxes') {
        // Find all empty boxes
        const emptyBoxes = [];
        const cols = Math.ceil(gridCanvas.width / BOX_SIZE);
        const rows = Math.ceil(gridCanvas.height / BOX_SIZE);

        for (let x = 0; x < cols; x++) {
            for (let y = 0; y < rows; y++) {
                const key = `${x},${y}`;
                if (!boxes[key]) {
                    emptyBoxes.push({
                        x: x * BOX_SIZE + BOX_SIZE / 2,
                        y: y * BOX_SIZE + BOX_SIZE / 2
                    });
                }
            }
        }

        // Launch rockets at all empty boxes
        emptyBoxes.forEach(box => {
            rockets.push(new Rocket(box.x, box.y));
        });

        // Change to reset mode
        fillAllBtn.textContent = 'Reset';
    } else {
        // Reset everything
        boxes = {};
        clickCount = 0;
        rockets = [];
        particles = [];
        fillAllBtn.style.display = 'none';
        fillAllBtn.textContent = 'Fill All Boxes';
    }
});

setInterval(() => {
    if (activeEffect === 'fireworks') {
        let x;
        if (Math.random() < 0.5) {
            x = Math.random() * (fireworksCanvas.width * 0.3);
        } else {
            x = (fireworksCanvas.width * 0.7) + Math.random() * (fireworksCanvas.width * 0.3);
        }

        const y = Math.random() * (fireworksCanvas.height * 0.6) + (fireworksCanvas.height * 0.1);
        rockets.push(new Rocket(x, y));
    }
}, 1000);

window.addEventListener('resize', () => {
    fireworksCanvas.width = window.innerWidth;
    fireworksCanvas.height = window.innerHeight;
});

updateEffects();
