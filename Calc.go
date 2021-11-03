package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Knetic/govaluate"
)

func showCalci() {
	output:=""
	input := widget.NewLabel(output)
	isHistory:=false
	historyStr:=""
	history:=widget.NewLabel(historyStr);
	var histarr []string;
	histtorybtn:=widget.NewButton("history",func() {
		if isHistory{
			historyStr=""
		}else{
			for i:=len(histarr)-1;i>=0;i--{
				historyStr=historyStr+histarr[i];
				historyStr+="\n"
			}
		}
		isHistory=!isHistory;
		history.SetText(historyStr);
	})
	backbtn:=widget.NewButton("back",func() {
		if len(output)>0{
			output=output[:len(output)-1];
			input.SetText(output);
		}
	})
	clearbtn:=widget.NewButton("clear",func() {
		output="";
		input.SetText(output);
	})
	openbtn:=widget.NewButton("(",func() {
		output=output+"(";
		input.SetText(output);
	})
	closebtn:=widget.NewButton(")",func() {
		output=output+")";
		input.SetText(output);
	})
	dividebtn:=widget.NewButton("/",func() {
		output=output+"/";
		input.SetText(output);
	})
	sevenbtn:=widget.NewButton("7",func() {
		output=output+"7";
		input.SetText(output);
	})
	eightbtn:=widget.NewButton("8",func() {
		output=output+"8";
		input.SetText(output);
	})
	ninebtn:=widget.NewButton("9",func() {
		output=output+"9";
		input.SetText(output);
	})
	multiplybtn:=widget.NewButton("*",func() {
		output=output+"*";
		input.SetText(output);
	})
	fourbtn:=widget.NewButton("4",func() {
		output=output+"4";
		input.SetText(output);
	})
	fivebtn:=widget.NewButton("5",func() {
		output=output+"5";
		input.SetText(output);
	})
	sixbtn:=widget.NewButton("6",func() {
		output=output+"6";
		input.SetText(output);
	})
	minusbtn:=widget.NewButton("-",func() {
		output=output+"-";
		input.SetText(output);
	})
	onebtn:=widget.NewButton("1",func() {
		output=output+"1";
		input.SetText(output);
	})
	twobtn:=widget.NewButton("2",func() {
		output=output+"2";
		input.SetText(output);
	})
	threebtn:=widget.NewButton("3",func() {
		output=output+"3";
		input.SetText(output);
	})
	plusbtn:=widget.NewButton("+",func() {
		output=output+"+";
		input.SetText(output);
	})
	zerobtn:=widget.NewButton("0",func() {
		output=output+"0";
		input.SetText(output);
	})
	dotbtn:=widget.NewButton(".",func() {
		output=output+".";
		input.SetText(output);
	})
	equalbtn:=widget.NewButton("=",func() {
		expression, err := govaluate.NewEvaluableExpression(output);
		if err==nil{
			result, err := expression.Evaluate(nil);
			if err==nil{
				ans:=strconv.FormatFloat(result.(float64),'f',-1,64);
				strtoAppend:=output+" = "+ans;
				histarr=append(histarr, strtoAppend);
				output=ans;
			}else{
				output="err";
			}
		}else{
			output="err"
		}
		input.SetText(output);
	})

	calciContainer:=container.NewVBox(container.NewVBox(
		input,
		history,
		container.NewGridWithColumns(1,
		container.NewGridWithColumns(2,
		histtorybtn,
		backbtn,
		),
		container.NewGridWithColumns(4,
		clearbtn,
		openbtn,
		closebtn,
		dividebtn,
		),
		container.NewGridWithColumns(4,
		sevenbtn,
		eightbtn,
		ninebtn,
		multiplybtn,
		),
		container.NewGridWithColumns(4,
		fourbtn,
		fivebtn,
		sixbtn,
		minusbtn,
		),
		container.NewGridWithColumns(4,
		onebtn,
		twobtn,
		threebtn,
		plusbtn,
		),
		container.NewGridWithColumns(2,
		container.NewGridWithColumns(2,
			zerobtn,
			dotbtn,
		   ),
		   equalbtn,
		),
		),
	))
	w:=myApp.NewWindow("Calc");
	w.Resize(fyne.NewSize(500,200));
	w.SetContent(
		container.NewBorder(deskBtn,nil,nil,nil,calciContainer),
	)
	w.Show()
}