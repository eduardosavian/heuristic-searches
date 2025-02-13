import pandas as pd
import matplotlib.pyplot as plt
import seaborn as sns

# Carregar os dados
df = pd.read_csv("hs.csv")

# Converter colunas categóricas para string para evitar erros
df["Heuristica"] = df["Heuristica"].astype(str)

# Criar uma figura com múltiplos gráficos
plt.figure(figsize=(25, 15))

# Boxplot do Valor por Heurística
plt.subplot(3, 3, 1)
sns.boxplot(x="Heuristica", y="Valor", data=df)
plt.xticks(rotation=45)
plt.title("Distribuição do Valor por Heurística")

# Gráfico de linha do Tempo vs Valor
plt.subplot(3, 3, 2)
sns.lineplot(x="Replicacao", y="Valor", hue="Heuristica", data=df, marker="o")
plt.title("Evolução do Valor por Replicação")

# Histograma do Tempo de Execução
plt.subplot(3, 3, 3)
sns.histplot(df["Tempo"], bins=20, kde=True)
plt.title("Distribuição do Tempo de Execução")

# Scatter plot de Iterações vs Valor
plt.subplot(3, 3, 4)
sns.scatterplot(x="Iteracoes", y="Valor", hue="Heuristica", data=df)
plt.title("Iterações vs Valor")

# Gráfico de barras do Valor médio por Heurística
plt.subplot(3, 3, 5)
sns.barplot(x="Heuristica", y="Valor", data=df, estimator=sum)
plt.xticks(rotation=45)
plt.title("Valor Total por Heurística")

# Heatmap de correlação entre variáveis numéricas
plt.subplot(3, 3, 6)
numeric_cols = df.select_dtypes(include=["number"])
sns.heatmap(numeric_cols.corr(), annot=True, cmap="coolwarm", fmt=".2f")
plt.title("Correlação entre Variáveis Numéricas")

# Gráfico de dispersão do Tempo vs Valor
plt.subplot(3, 3, 7)
sns.scatterplot(x="Tempo", y="Valor", hue="Heuristica", data=df)
plt.title("Dispersão do Tempo vs Valor")

# Boxplot do Tempo por Heurística
plt.subplot(3, 3, 8)
sns.boxplot(x="Heuristica", y="Tempo", data=df)
plt.xticks(rotation=45)
plt.title("Distribuição do Tempo por Heurística")

# Gráfico de barras do Número de Iterações médio por Heurística
plt.subplot(3, 3, 9)
sns.barplot(x="Heuristica", y="Iteracoes", data=df, estimator=sum)
plt.xticks(rotation=45)
plt.title("Número Total de Iterações por Heurística")

# Ajustar layout e mostrar
plt.tight_layout()
plt.show()
