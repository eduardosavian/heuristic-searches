import pandas as pd
import matplotlib.pyplot as plt

# Leia os dados do arquivo
data = pd.read_csv('mlsb.csv')

# Plote os dados
plt.figure(figsize=(10, 6))
for key, grp in data.groupby(['Replicacao']):
    plt.plot(grp['Iteracoes'], grp['Valor'], label=f'Replicacao {key}')

plt.xlabel('Iteracoes')
plt.ylabel('Valor')
plt.title('Iteracoes vs Valor for different Replicacoes')
plt.legend()
plt.grid(True)
plt.show()