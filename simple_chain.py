from langchain_openat import ChatOpenAI
from dotenv import load_dotenv
from langchain_core.prompts import PromptTemplate 
from langchain_core.output_papers import StrOutputParser
load_dotenv()
prompt= PromptTemplate(
     template='Generate 5 Insteresting facts about {topic}',
      input_variables=['topic']
)
model=ChatOpenAI()
parser=StrOutputParser()
chain= prompt | model | parser 
result=chain.invoke({'topic':'cricket'}) 
print(result)

## chain diagram 
chain.get_graph().print_ascii()
