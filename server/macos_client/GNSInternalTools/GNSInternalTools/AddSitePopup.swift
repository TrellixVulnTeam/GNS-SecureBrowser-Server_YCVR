//
//  AddSitePopup.swift
//  GNSInternalTools
//
//  Created by Peter Pham on 2/26/22.
//

import SwiftUI

struct AddSitePopup: View {
    var title: String
    var message: String
    var buttonText: String
    @Binding var show: Bool
    @Binding var freesiteTable: [UInt32]
    @State var selectedIdx: UInt32
    @State var username: String
    @State var password: String
    @State var code: String
    @Binding var client: GNSRPC_GNSBadgeDataClient?
    
    var body: some View {
        ZStack {
            if show {
                // PopUp background color
                Color.black.opacity(show ? 0.3 : 0).edgesIgnoringSafeArea(.all)
                
                // PopUp Window
                VStack(alignment: .center, spacing: 0) {
                    Text(title)
                        .frame(maxWidth: .infinity)
                        .frame(height: 45, alignment: .center)
                        .font(Font.system(size: 23, weight: .semibold))
                        .foregroundColor(Color.white)
                        .background(Color(#colorLiteral(red: 0.6196078431, green: 0.1098039216, blue: 0.2509803922, alpha: 1)))
                    
                    VStack {
                        
                        Picker("Idx:", selection: $selectedIdx) {
                                ForEach(freesiteTable, id: \.self) { a in
                                    Text(String(a))
                                }
                        }
                        .pickerStyle(.menu)
                        .padding()
                        TextField("Code", text: $code)
                        TextField("Username", text: $username)
                        TextField("Password", text: $password)
                    }
                    
                    Button(action: {
                        // Dismiss the PopUp
                        withAnimation(.linear(duration: 0.3)) {
                            show = false
                        }
                    }, label: {
                        Text("Cancel")
                            .frame(maxWidth: .infinity)
                            .frame(height: 54, alignment: .center)
                            .foregroundColor(Color.white)
                            .background(Color(#colorLiteral(red: 0.6196078431, green: 0.1098039216, blue: 0.2509803922, alpha: 1)))
                            .font(Font.system(size: 23, weight: .semibold))
                    }).buttonStyle(PlainButtonStyle())
                   
                    Button(action: {
                        // Dismiss the PopUp
                        var newsite = GNSRPC_SiteCred.init()
                        newsite.code = code
                        newsite.username = username
                        newsite .password = password
                        newsite.idx = selectedIdx
                        if(newsite.username.count > 0) {
                            let result = api().writeSite(client: client!, data: newsite)
                            result.whenSuccess { a in
                                print("Finish saving site idx:%d with username: %s",newsite.idx, newsite.username)
                            }
                        }
                        withAnimation(.linear(duration: 0.3)) {
                            show = false
                        }
                    }, label: {
                        Text(buttonText)
                            .frame(maxWidth: .infinity)
                            .frame(height: 54, alignment: .center)
                            .foregroundColor(Color.white)
                            .background(Color(#colorLiteral(red: 0.6196078431, green: 0.1098039216, blue: 0.2509803922, alpha: 1)))
                            .font(Font.system(size: 23, weight: .semibold))
                    }).buttonStyle(PlainButtonStyle())
                }
                .frame(maxWidth: 500)
                .border(Color.white, width: 2)
                .background(Color(#colorLiteral(red: 0.737254902, green: 0.1294117647, blue: 0.2941176471, alpha: 1)))
            }
        }
    }
}
/*
 struct AddSitePopup_Previews: PreviewProvider {
 static var previews: some View {
 AddSitePopup()
 }
 }*/
