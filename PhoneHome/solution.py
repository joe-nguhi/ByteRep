# SpaceWeekDay3 - Phone Home 
# Determine the time it takes to send a message from your space ship to your planet.
SPEED_OF_LIGHT = 300_000
SAT_DELAY = 0.5
def send_message(route):
    total_distance = 0
    for dst in route:
        total_distance += dst

    n_sats = len(route) - 1
    
    return round(total_distance/SPEED_OF_LIGHT + n_sats*SAT_DELAY, 4)


print(send_message([300_000, 300_000]))
print(send_message([384_400, 384_400]))
print(send_message([1_000_000,500_000_000,1_000_000]))