# only best for real pdf not scanned one 
from langchain_community.documents_loader import PyPDFLoader 
loader = PyPDFLoader('sample.pdf')
doc=loader.load()
print(len(docs))

# what is we want to load a whole folder how wouuld you do it 
# here is how 
from langchain_comunity.documents_loader import DirectoryLoader 
loader2=DirectoryLoader( path="books", glob="*.pdf",  loader_cls=PyPDFLoader)
doc2=loader.load()
print(len(docs))
# this is very compuntaion heavy f
# so in real life we use lazy loading 
#   lazy_load()
