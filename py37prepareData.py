import pandas as pd 

df = pd.read_csv("data/hcvdat0.csv")

# df
# Numerical & No missing values & Head 

print(df.dtypes) 

df["Sex"].unique()

gender_dict = {'f':0,'m':1}

print(df["Sex"].map(gender_dict))

df["Sex"] = df["Sex"].map(gender_dict)

print(df.head())

print(df["Category"].value_counts())

print(df["Category"].unique())

hcdict = {
    '0=Blood Donor':0,
    '0s=suspect Blood Donor':1,
    '1=Hepatitis':1,
    '2=Fibrosis':1,
    '3=Cirrhosis':1,
}

df["Target"] = df["Category"].map(hcdict)

print(df.head())

print(df.isna())

print(df.isna().sum())

df.fillna(0)

df.fillna(0, inplace=True)
 
print(df.isna().sum())

print(df.dtypes)

# Split Train/Test
print(df.head)

print(df.shape)

print(df.columns)

df = df[[ 'Age', 'Sex', 'ALB', 'ALP', 'ALT', 'AST',
       'BIL', 'CHE', 'CHOL', 'CREA', 'GGT', 'PROT', 'Target']]

print(df)

train = df.sample(frac=0.7, random_state=200)
test = df.drop(train.index)

print(train.head())
print(test.head())

train.to_csv("data/prepdata/trainhcvData.csv", index=False, header=False)
test.to_csv("data/prepdata/testhcvData.csv", index=False, header=False)
df.to_csv("data/cleanhcvData.csv")

print(test.head())

print(test.iloc[0])
print(test.iloc[0].tolist())

print(test.head())

print(test.iloc[183])
print(test.iloc[183].tolist())
