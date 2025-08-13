# simple qa agent 
import os 
#import dotenv import load_dotenv
from langchain_openai import ChatOpenAI
from langchain.prompts import PromptTemplate
llm= ChatOpenAI(model="your_model_name",max_tokens=1000,temperature=0)
template= """
You are a helpful ai assistant . Your task is to answe the user question to best of your 
ability 

user's question:{question}

please provide a clear and consise answer:
"""
prompt=PromptTemplate(template=template,input_variables=["question"])
# make a chain 
qa_chain= prompt|llm 

def get_answer(question):
    """_summary_
 get an answer to the given question using the QA chain 
    Args:
        question (_type_): _description_
    """
    input_variables={"question":question}
    response=qa_chain.invoke(input_variables).contents
    return response 
## useage 
question="What is the capital of france"
answer= get_answer(question)
print(f"Question:{question}")
print(f"Answer:{answer}")