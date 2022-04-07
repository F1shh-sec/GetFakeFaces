from DuplicateRemover import DuplicateRemover

dirname = "faces"

# Remove Duplicates
dr = DuplicateRemover(dirname)
dr.find_duplicates()
