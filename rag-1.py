# document loaded 

# text loaded most simple when we use it when we want to load chat or simple text 
from langchain_community.document_loaders import TextLoader 
from langchain_openai import ChatOpenAI
from langchain_core.output_parsers import StrOutputParser
from langchain_core.propmts import PromptTemplate 
from dotenv import load_dotenv 
load_dotenv()
model=ChatOpenAI()
prompt = PromptTemplate(
     template='Generate a 3 points summray from this given text \n{text}',
     input_varaiables=['text']
)
parser=StrOutputParser()


loader=TextLoader("sample.txt",encoding='utf-8' )
docs =loader.load()
chain= prompt | model | parser
print(chain.invoke({'text':docs[0].page_content}))
# loaded in list data data type 

# two parameters in this 
