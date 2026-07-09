namespace targets {

    class Alien {
        public:
            Alien(int x, int y) {
                x_coordinate = x;
                y_coordinate = y;
            }
    
            int x_coordinate;
            int y_coordinate;

            int get_health() {
                return health_level;
            }

            bool hit() {
                if(health_level > 0) {
                    health_level--;
                }
                
                return true;
            }

            bool is_alive() {
                if (health_level <= 0) {
                    return false;
                }
                return true;
            }

            bool teleport(int x_new, int y_new) {
                x_coordinate = x_new;
                y_coordinate = y_new;
                return true;
            }


            bool collision_detection(Alien l) {
                if(l.x_coordinate == x_coordinate && l.y_coordinate == y_coordinate){
                    return true;
                }
                return false;
            }

    
        private:
            int health_level{3};
    
    };

    
}  // namespace targets
