# Fireworks Effect Technical Analysis

## Overview
The fireworks effect combines an interactive grid system with particle-based fireworks animation. Users click to launch rockets that explode into colorful particles while filling grid boxes, creating a gamified visual experience.

## Technical Implementation

### Multi-Canvas Architecture
- **gridCanvas**: Static grid background with fillable boxes
- **fireworksCanvas**: Dynamic particle animation layer
- **Layering**: Grid behind, fireworks on top for proper visual hierarchy

### Grid System

#### Configuration
```javascript
BOX_SIZE = 40        // Grid cell dimensions (40×40 pixels)
boxes = {}           // Hash map for filled boxes
```

#### Grid Generation
- **Calculation**: `cols = ceil(width/40)`, `rows = ceil(height/40)`
- **Rendering**: Nested loops draw grid lines with `strokeRect()`
- **Storage**: Box states stored as `"x,y": color` key-value pairs

#### Box Filling Algorithm
```javascript
boxX = Math.floor(x / BOX_SIZE)
boxY = Math.floor(y / BOX_SIZE)
key = `${boxX},${boxY}`
boxes[key] = color
```

### Rocket Class Architecture

#### Properties
- `x, y`: Current position coordinates
- `targetY`: Destination height for explosion
- `color`: HSL color (`hsl(${Math.random() * 360}, 100%, 50%)`)
- `speed`: Vertical velocity (12 pixels/frame)
- `wobble`: Horizontal oscillation factor

#### Physics Simulation
```javascript
this.y -= this.speed                    // Constant upward velocity
this.x += Math.sin(this.y * 0.1) * 0.5  // Sine wave wobble effect
```

#### Explosion Trigger
- Condition: `this.y <= this.targetY`
- Actions: Remove rocket, fill grid box, create particle explosion

### Particle Class Architecture

#### Properties
- `x, y`: Position coordinates
- `vx, vy`: Velocity components
- `color`: Inherited from parent rocket
- `alpha`: Opacity (1.0 to 0.0)
- `friction`: Velocity decay (0.98)
- `gravity`: Downward acceleration (0.05)

#### Physics Engine
```javascript
this.vx *= this.friction    // Air resistance
this.vy *= this.friction
this.vy += this.gravity     // Gravitational pull
this.x += this.vx          // Position integration
this.y += this.vy
this.alpha -= 0.015        // Fade out over time
```

#### Explosion Generation
```javascript
const angle = Math.random() * Math.PI * 2      // Random direction
const velocity = Math.random() * 5 + 2         // Random speed (2-7)
this.vx = Math.cos(angle) * velocity           // X component
this.vy = Math.sin(angle) * velocity           // Y component
```

## Animation Workflow

### Frame Processing
1. **Trail Effect**: Apply fade overlay (`rgba(0,0,0,0.2)`)
2. **Rocket Updates**: Move rockets, check collision, trigger explosions
3. **Particle Updates**: Apply physics, fade particles, remove dead ones
4. **Rendering**: Draw all active rockets and particles
5. **Frame Request**: Continue animation loop

### Composite Operations
```javascript
fireworksCtx.globalCompositeOperation = 'destination-out'  // Fade effect
fireworksCtx.globalCompositeOperation = 'source-over'      // Normal drawing
```

## Interactive Features

### Click-to-Launch System
- **Event**: `window.addEventListener('click')`
- **Action**: Create new rocket at click coordinates
- **Counter**: Track clicks for special features

### Auto-Launch System
```javascript
setInterval(() => {
    // Launch from sides (30% left or 30% right)
    // Target random height (10-70% of screen)
}, 1000)
```

### Fill All Boxes Feature
- **Trigger**: 20+ user clicks
- **Algorithm**: Find empty boxes, launch rockets at each
- **State Management**: Toggle between "Fill All" and "Reset" modes

## Mathematical Analysis

### Particle Physics
- **Initial Velocity**: Polar coordinates converted to Cartesian
- **Trajectory**: Parabolic motion under gravity
- **Decay**: Exponential velocity reduction via friction

### Color System
- **Generation**: HSL with random hue (0-360°)
- **Saturation**: 100% for vibrant colors
- **Lightness**: 50% for optimal visibility

### Performance Optimization
- **Object Pooling**: Particles removed when alpha ≤ 0
- **Efficient Rendering**: Single draw call per particle
- **Memory Management**: Dynamic array resizing

## Visual Effects

### Trail System
- **Method**: Overlay semi-transparent black rectangle
- **Opacity**: 20% per frame creates smooth trails
- **Result**: Particles leave fading streaks

### Explosion Dynamics
- **Particle Count**: 100 particles per explosion
- **Distribution**: Uniform radial spread
- **Lifespan**: ~67 frames (alpha 1.0 → 0.0 at -0.015/frame)

### Grid Integration
- **Synchronization**: Rocket explosions fill corresponding grid boxes
- **Visual Feedback**: Immediate color application
- **Persistence**: Filled boxes remain until reset

## Browser Compatibility
- **Canvas 2D**: Universal support
- **Event Handling**: Standard DOM events
- **Performance**: Optimized for 60 FPS
- **Memory**: Efficient particle lifecycle management