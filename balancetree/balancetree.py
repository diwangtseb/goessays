class TreeNode:
    def __init__(self, val: int, left=None, right=None) -> None:
        self.val = val
        self.left = left
        self.right = right

    def __repr__(self) -> str:
        return f"val: {self.val}, left: {self.left}, right: {self.right}"

    def __str__(self) -> str:
        return str(self.val)

def to_binary_tree(items)->TreeNode:
    n = len(items)
    if n == 0 :
        return None
    def inner(index:int = 0) -> TreeNode:
        if n <= index or items[index] is None:
            return None
        node = TreeNode(items[index])
        node.left = inner(2 * index + 1)
        node.right = inner(2 * index + 2)
        return node
    return inner()

if __name__ == "__main__":
    root = to_binary_tree([1, 2, 3, None, None, 4, 5])
    print(root.__repr__())