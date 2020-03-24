#include <iostream>
#include <ctime>
#include <cstdlib>

using namespace std;

class BST
{
	int data;
	BST *left, *right;

public:
	// Default constructor.
	BST();

	// Parameterized constructor.
	BST(int);

	// Insert function.
	BST *Insert(BST *, int);

	// Inorder traversal.
	void Inorder(BST *);

	// Get minimum value node.
	BST *MinValueNode(BST *);

	// Delete a node.
	BST *Delete(BST *, int);
};

// Default Constructor definition.
BST ::BST() : data(0), left(NULL), right(NULL) {}

// Parameterized Constructor definition.
BST ::BST(int value)
{
	data = value;
	left = right = NULL;
}

// Insert function definition.
BST *BST ::Insert(BST *root, int value)
{
	if (!root)
	{
		// Insert the first node, if root is NULL.
		return new BST(value);
	}

	// Insert data.
	if (value > root->data)
	{
		// Insert right node data, if the 'value'
		// to be inserted is greater than 'root' node data.

		// Process right nodes.
		root->right = Insert(root->right, value);
	}
	else
	{
		// Insert left node data, if the 'value'
		// to be inserted is greater than 'root' node data.

		// Process left nodes.
		root->left = Insert(root->left, value);
	}

	// Return 'root' node, after insertion.
	return root;
}

// Inorder traversal function.
// This gives data in sorted order.
void BST ::Inorder(BST *root)
{
	if (!root)
	{
		return;
	}
	Inorder(root->left);
	cout << root->data << endl;
	Inorder(root->right);
}

BST *BST ::MinValueNode(BST *root) 
{ 
    BST *current = root; 
  
    /* loop down to find the leftmost leaf */
    while (current && current->left != NULL) 
        current = current->left; 
  
    return current; 
} 

BST *BST ::Delete(BST *root, int key)
{
	// base case
	if (root == NULL)
		return root;

	// If the key to be deleted is smaller than the root's key,
	// then it lies in left subtree
	if (key < root->data)
		root->left = Delete(root->left, key);

	// If the key to be deleted is greater than the root's key,
	// then it lies in right subtree
	else if (key > root->data)
		root->right = Delete(root->right, key);

	// if key is same as root's key, then This is the node
	// to be deleted
	else
	{
		// node with only one child or no child
		if (root->left == NULL)
		{
			BST *temp = root->right;
			free(root);
			return temp;
		}
		else if (root->right == NULL)
		{
			BST *temp = root->left;
			free(root);
			return temp;
		}

		// node with two children: Get the inorder successor (smallest
		// in the right subtree)
		BST *temp = MinValueNode(root->right);

		// Copy the inorder successor's content to this node
		root->data = temp->data;

		// Delete the inorder successor
		root->right = Delete(root->right, temp->data);
	}
	return root;
}

int main(int argc, char const *argv[])
{
	srand(time(NULL));

	BST b, *root = NULL;

	int S = 1000000;
	int n = 1000000;
	if (argc >= 2)
	{
		S = atoi(argv[1]);

		if (argc == 3)
		{
			n = atoi(argv[2]);
		}
	}

	for (int i = 0; i < S; i++)
	{
		if (i == 0)
		{
			root = b.Insert(root, rand());
		}
		else
		{
			b.Insert(root, rand());
		}
	}

	for (int i = 0; i < n; i++)
	{
		root = b.Delete(root, rand());
	}

	return 0;
}