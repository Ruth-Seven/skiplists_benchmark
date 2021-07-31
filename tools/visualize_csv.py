#!/usr/bin/python3
# coding: utf-8
import argparse
from pathlib import Path

import pandas as pd
import numpy as np
import seaborn as sns
import matplotlib.pyplot as plt


reg = {
    ".*RandomInserts":"RandomInserts.png",
    "^(?!.*Worst)(?!.*Random).*Inserts$": "Insert.png",
    "^.*WorstInserts$": "WorstInserts.png",
    
    ".*AvgSearch": "AvgSearch.png",
    ".*SearchEnd": "SearchEnd.png",
   
    "^(?!.*Worst)(?!.*Random).*Delete$": "Delete.png",
    ".*WorstDelete": "WorstDelete.png",
    ".*RandomDelete": "RandomDelete.png",
}
folder = "./result"

def plotAndSave(df, regex=".*", save_name=".png"): 
    data = df.filter(regex=regex)
    sns.set_theme()
    fig = sns.relplot(
        data=data, kind="line",
    )
    fig.savefig( Path(folder) / save_name)

def render_md(reg={}):
    fo = open(Path(folder) / "./png.md","w")
    fo.write("## Analysis Result Pns\n\n")
    for key in reg: 
        fo.write("#### {:s}\n\n".format(reg[key]))
        fo.write("![{:s}](./{:s})\n\n".format(reg[key], reg[key]))
    fo.close()


def read_df():
	plt.rcParams['figure.dpi'] = 300
	plt.rcParams['savefig.dpi'] = 300
	p = Path(folder) / "./result.csv"
	df = pd.read_csv(p, header=None).dropna(axis=1).T
	df.columns = df.iloc[0] 
	df.drop(index=0, inplace=True)
	
	df = df.set_index(df['iterations']).drop(columns="iterations")
	return df


if '__main__' == __name__:
	parser = argparse.ArgumentParser(description='Process plotting and rending.')
	parser.add_argument('-folder', dest="folder", type=str,
						help=' the folder which solved')
	folder = parser.parse_args().folder

	df = read_df()
	render_md(reg)
	# you can use "@@include[my-file.md](includes/my-file.md)" to include the png.md
	for key in reg:
		plotAndSave(df, key, reg[key])


