const dis = () => {

    const pixel_density = window.devicePixelRatio || 1;
    const optimal_dpi = pixel_density * 96;
    const optimal_width = window.innerWidth * pixel_density;
    const optimal_height = window.innerHeight * pixel_density;
    
    const width = Math.floor(optimal_width);
    const height = Math.floor(optimal_height);
    const dpi = Math.floor(optimal_dpi);
    
    const arr = [width, height, dpi];
    
      return arr;
    };
    
export default dis;
    