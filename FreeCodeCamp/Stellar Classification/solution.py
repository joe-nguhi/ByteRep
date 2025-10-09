def classification(temp):
    match temp:
        case _ if temp >= 30000:
            return ("O")
        case _ if temp >= 10000:
            return "B"
        case _ if temp >= 7500:
            return "A"
        case _ if temp >= 6000:
            return "F"
        case _ if temp >= 5200:
            return "G"
        case _ if temp >= 3700:
            return "K"
        case _ if temp >= 0:
            return "M"
        case _:
            return "Unclassified"
    

print(f"Temp:\t5778  Star Class:{classification(5778)}")
print(f"Temp:\t2400  Star Class:{classification(2400)}")
print(f"Temp:\t9999  Star Class:{classification(9999)}")
