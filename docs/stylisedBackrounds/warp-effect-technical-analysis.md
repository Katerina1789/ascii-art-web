# Warp Effect Technical Analysis

## Overview
The warp effect creates a dynamic starfield animation that simulates traveling through space at high speed. Stars appear to move from the center outward, creating a tunnel-like visual effect.

## Technical Implementation

### Canvas Setup
- **Element**: `warpCanvas` with 2D rendering context
- **Dimensions**: Full viewport (`window.innerWidth` × `window.innerHeight`)
- **Responsive**: Auto-resizes on window resize events

### Core Parameters
```javascript
STAR_COUNT = 400    // Total number of stars
SPEED = 0.1         // Base movement speed multiplier
STAR_SIZE = 2       // Base star size
SPREAD = 800        // 3D space dimensions
```

### Star Class Architecture

#### Properties
- `x, y`: 2D coordinates in 3D space (-SPREAD to +SPREAD)
- `z`: Depth coordinate (0 to SPREAD)
- `pz`: Previous z-position for trail effects

#### Methods

**reset()**: Initializes star at random 3D position
```javascript
this.x = (Math.random() - 0.5) * SPREAD * 2  // -800 to 800
this.y = (Math.random() - 0.5) * SPREAD * 2  // -800 to 800
this.z = Math.random() * SPREAD              // 0 to 800
```

**update()**: Moves star toward viewer
```javascript
this.z = this.z - (SPEED * 20)  // Decreases z by 2 units per frame
```
- Stars reset when `z <= 0` (reached viewer)

**draw()**: Projects 3D coordinates to 2D screen
```javascript
sx = (this.x / this.z) * 100 + warpCenterX  // Perspective projection
sy = (this.y / this.z) * 100 + warpCenterY
size = (1 - this.z / SPREAD) * STAR_SIZE * 3 // Size increases as z decreases
```

## Mathematical Analysis

### Perspective Projection
The effect uses perspective division to create depth illusion:
- **Formula**: `screen_pos = (world_pos / depth) * focal_length + center`
- **Focal Length**: 100 pixels
- **Result**: Objects closer to viewer (smaller z) appear larger and more spread out

### Opacity Calculation
```javascript
opacity = (1 - this.z / SPREAD)
```
- Stars fade in as they approach (z decreases)
- Maximum opacity when z = 0
- Invisible when z = SPREAD

### Size Scaling
```javascript
size = (1 - this.z / SPREAD) * STAR_SIZE * 3
```
- Stars grow from 0 to 6 pixels as they approach
- Creates illusion of acceleration

## Animation Workflow

### Frame Cycle
1. **Clear Canvas**: Fill with black background
2. **Update Stars**: Move each star closer (decrease z)
3. **Reset Distant Stars**: Reinitialize stars that reached viewer
4. **Draw Stars**: Render each star with calculated position/size/opacity
5. **Request Next Frame**: Continue animation loop

### Performance Characteristics
- **Frame Rate**: 60 FPS (requestAnimationFrame)
- **Calculations per Frame**: 400 stars × 4 operations = 1,600 operations
- **Memory Usage**: Minimal (400 star objects)

## Visual Effects

### Tunnel Effect
- Stars appear to emerge from center point
- Radial movement creates depth perception
- Faster movement near edges enhances perspective

### Speed Illusion
- Constant z-velocity creates acceleration appearance
- Size/opacity changes reinforce speed sensation
- Black background emphasizes star trails

### Immersion Factors
- Full-screen coverage
- Smooth 60 FPS animation
- Realistic perspective mathematics
- Seamless star recycling

## Browser Compatibility
- **Canvas 2D**: Universal support
- **requestAnimationFrame**: IE10+
- **Performance**: Optimized for modern browsers
- **Fallback**: Graceful degradation on older systems