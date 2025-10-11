THRESHOLD = 0.8


def has_exoplanet(readings) -> bool:
    """
    Given a list of readings from a star, return True if there is a chance
    that the star has an exoplanet.
    """
    parsed = [from_base_36(r) for r in readings]

    average = get_average(parsed)

    for r in parsed:
        if r is None:
            continue
        if r <= average * THRESHOLD:
            return True

    return False


def get_average(readings) -> float:
    s = 0

    for r in readings:
        if r is None:
            continue
        s += r

    return s / len(readings)


def from_base_36(s) -> int | None:
    d = ord(s)

    if ord('0') <= d <= ord('9'):
        return d - ord('0')

    if ord('A') <= d <= ord('Z'):
        return d - ord('A') + 10

    return None


# {"665544554", false},
# {"FGFFCFFGG", true},
# {"MONOPLONOMONPLNOMPNOMP", false},
# {"FREECODECAMP", true},
# {"9AB98AB9BC98A", false},
# {"ZXXWYZXYWYXZEGZXWYZXYGEE", true},

print(has_exoplanet("665544554"))
print(has_exoplanet("FGFFCFFGG"))
print(has_exoplanet("MONOPLONOMONPLNOMPNOMP"))
print(has_exoplanet("FREECODECAMP"))
print(has_exoplanet("9AB98AB9BC98A"))
print(has_exoplanet("ZXXWYZXYWYXZEGZXWYZXYGEE"))
