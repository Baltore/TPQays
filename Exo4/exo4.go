package exo

func Ft_profit(prices []int) int {

	// maxbenef = 0
	// minPrix  = prices[0]

	// Prend le premier indice de prices
	// verifie si a partir de l'indice de prices si il y a un prices inferieur sinon
	// on fait minPrix - price = profit pour chaque element du tableau, pour stock
	// le meilleur profit et si il est supérieur à l'ancien maxbenef on le stock et
	// on le retourne

	maxbenef := 0
	minPrix := prices[0]

	for _, price := range prices[1:] {
		if price < minPrix {
			minPrix = price
		} else {
			profit := price - minPrix
			if profit > maxbenef {
				maxbenef = profit
			}
		}
	}

	return maxbenef
}
