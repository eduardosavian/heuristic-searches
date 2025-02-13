import pandas as pd
import matplotlib.pyplot as plt
import seaborn as sns

# Carregar o arquivo CSV
df = pd.read_csv("nmlrs.csv")

# Criar gráficos
sns.set(style="whitegrid")

# 1. Evolução do Valor ao longo das Replicações
plt.figure(figsize=(8, 5))
sns.lineplot(x="Replicacao", y="Valor", data=df, marker="o")
plt.title("Evolução do Valor ao longo das Replicações")
plt.xlabel("Replicação")
plt.ylabel("Valor")
plt.show()

# 2. Distribuição do Tempo de Execução
plt.figure(figsize=(8, 5))
sns.histplot(df["Tempo"], bins=20, kde=True)
plt.title("Distribuição do Tempo de Execução")
plt.xlabel("Tempo")
plt.ylabel("Frequência")
plt.show()

# 3. Correlação entre Tempo e Valor
plt.figure(figsize=(8, 5))
sns.scatterplot(x="Tempo", y="Valor", data=df)
plt.title("Correlação entre Tempo e Valor")
plt.xlabel("Tempo")
plt.ylabel("Valor")
plt.show()

# 4. Boxplot do Valor por Replicação
plt.figure(figsize=(8, 5))
sns.boxplot(x="Replicacao", y="Valor", data=df)
plt.title("Boxplot do Valor por Replicação")
plt.xlabel("Replicação")
plt.ylabel("Valor")
plt.show()

# 5. Histograma do Valor
plt.figure(figsize=(8, 5))
sns.histplot(df["Valor"], bins=20, kde=True)
plt.title("Distribuição do Valor")
plt.xlabel("Valor")
plt.ylabel("Frequência")
plt.show()

# 6. Gráfico de barras da média do Valor por Replicação
plt.figure(figsize=(8, 5))
sns.barplot(x="Replicacao", y="Valor", data=df, estimator=sum, ci=None)
plt.title("Soma do Valor por Replicação")
plt.xlabel("Replicação")
plt.ylabel("Soma do Valor")
plt.show()
