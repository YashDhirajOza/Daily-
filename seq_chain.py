from langchain_openai import ChatOpenAI 
from dotenv import load_dotenv
from langchain_core,prompts import PromptTemplate 
from langchain_core.output_parsers import StrOutputPParser
# loading the apis 
load_dotenv()
prom1= PromptTemplate(template= 'Generate a detailed report on {topic}',   input_variables=['topic'] )
prom2=PromptTemplate(template=' Generate a 5 point summary from the following text\n {text}',input_variables=['text'])
model=ChatOpenAI()
parser=StrOutputParser()
chain= prom1 | model | parser| model | parser
result= chain.invoke({'topic':'football'})
print(result)
chain.get_graph().print_ascii()
