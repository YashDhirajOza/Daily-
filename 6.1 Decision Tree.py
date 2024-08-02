from sklearn.datasets import load_iris
from sklearn.tree import DecisionTreeClassifier
from sklearn.tree import export_graphviz
from graphviz import Source 

# Load the iris dataset with feature names
iris = load_iris(as_frame=True)

# Check the actual feature names in the dataset
print(iris.feature_names)

# Select the appropriate columns
X_iris = iris.data[["petal length (cm)", "petal width (cm)"]].values
y_iris = iris.target

# Create and train the decision tree classifier
tree_clf = DecisionTreeClassifier(max_depth=2, random_state=42)
tree_clf.fit(X_iris, y_iris)

tree_clf.predict_proba([[5,1.5]]).round(3)

tree_clf.predict([[5,1.5]])

