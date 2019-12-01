def get_fuel(mass: int):
    fuel = mass//3 - 2

    if fuel <= 0:
        return 0    
    return fuel + get_fuel(fuel)

if __name__ == "__main__":
    f = open("1.in")
    lines = f.readlines()

    fuel = 0
    for line in lines:
        mass = int(line.rstrip())
        fuel += get_fuel(mass)

    print(fuel)
