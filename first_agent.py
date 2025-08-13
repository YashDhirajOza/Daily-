from langchain_openai import ChatOpenAI
from langchain_core.runnables.history import RunnableWithMessageHistory
from langchain.memory import ChatMessageHistory
from langchain_core.prompts import ChatPromptTemplate,MessagesPlaceholder

import os  
## 
## leaving the part where you really load the api 
# from dotnev 
llm=ChatOpenAI(model="gtp-4o-mini",max_tokens=1000,temperature=0)
store ={}
# we used here a dict 
## simple in memory store for chat histroy 
def get_chat_history(session_id:str):
    if session_id not in store:
        store[session_id]=ChatMessageHistory()
    return store[session_id]

# propmt template 
prompt=ChatPromptTemplate.from_messages(
    [("system","you are a helpfull ai assistant"),
      MessagesPlaceholder(variable_name="history"),
      ("human","{input}")
    ]
)
# runnable chain 
chain = prompt| llm

# wrap the chain with message histry 
chain_with_history = RunnableWithMessageHistory(
    chain,
    get_chat_history,
    input_messages_key="input",
    history_messages_key="history"
)
# uses open ai 
# a simple in memeory store per seesion chat histories using chat message history 
#  chat pro