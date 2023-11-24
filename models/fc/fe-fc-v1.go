package fc

type Fc struct {
	Schema     string `json:"$schema"`
	Title      string `json:"title"`
	Type       string `json:"type"`
	Properties struct {
		Identificacion struct {
			Description string `json:"description"`
			Type        string `json:"type"`
			Properties  struct {
				Version struct {
					Type        string `json:"type"`
					Description string `json:"description"`
					Const       int    `json:"const"`
				} `json:"version"`
				Ambiente struct {
					Type        string   `json:"type"`
					Description string   `json:"description"`
					Enum        []string `json:"enum"`
				} `json:"ambiente"`
				TipoDte struct {
					Type        string `json:"type"`
					Description string `json:"description"`
					Const       string `json:"const"`
				} `json:"tipoDte"`
				NumeroControl struct {
					Type        string `json:"type"`
					Description string `json:"description"`
					MaxLength   int    `json:"maxLength"`
					MinLength   int    `json:"minLength"`
					Pattern     string `json:"pattern"`
				} `json:"numeroControl"`
				CodigoGeneracion struct {
					Type        string `json:"type"`
					Description string `json:"description"`
					MaxLength   int    `json:"maxLength"`
					MinLength   int    `json:"minLength"`
					Pattern     string `json:"pattern"`
				} `json:"codigoGeneracion"`
				TipoModelo struct {
					Type        string `json:"type"`
					Description string `json:"description"`
					Enum        []int  `json:"enum"`
				} `json:"tipoModelo"`
				TipoOperacion struct {
					Type        string `json:"type"`
					Description string `json:"description"`
					Enum        []int  `json:"enum"`
				} `json:"tipoOperacion"`
				TipoContingencia struct {
					Description string   `json:"description"`
					Type        []string `json:"type"`
					Enum        []any    `json:"enum"`
				} `json:"tipoContingencia"`
				MotivoContin struct {
					Description string   `json:"description"`
					Type        []string `json:"type"`
					MaxLength   int      `json:"maxLength"`
					MinLength   int      `json:"minLength"`
				} `json:"motivoContin"`
				FecEmi struct {
					Type        string `json:"type"`
					Description string `json:"description"`
					Format      string `json:"format"`
				} `json:"fecEmi"`
				HorEmi struct {
					Type        string `json:"type"`
					Description string `json:"description"`
					Pattern     string `json:"pattern"`
				} `json:"horEmi"`
				TipoMoneda struct {
					Type        string   `json:"type"`
					Description string   `json:"description"`
					Enum        []string `json:"enum"`
				} `json:"tipoMoneda"`
			} `json:"properties"`
			AllOf []struct {
				If struct {
					Properties struct {
						TipoOperacion struct {
							Const int `json:"const"`
						} `json:"tipoOperacion"`
					} `json:"properties"`
				} `json:"if"`
				Then struct {
					Properties struct {
						TipoModelo struct {
							Const int `json:"const"`
						} `json:"tipoModelo"`
						TipoContingencia struct {
							Type string `json:"type"`
						} `json:"tipoContingencia"`
						MotivoContin struct {
							Type string `json:"type"`
						} `json:"motivoContin"`
					} `json:"properties"`
				} `json:"then"`
				Else struct {
					Properties struct {
						TipoModelo struct {
							Enum []int `json:"enum"`
						} `json:"tipoModelo"`
					} `json:"properties"`
				} `json:"else,omitempty"`
			} `json:"allOf"`
			AdditionalProperties bool     `json:"additionalProperties"`
			Required             []string `json:"required"`
		} `json:"identificacion"`
		DocumentoRelacionado struct {
			Description string   `json:"description"`
			Type        []string `json:"type"`
			Items       struct {
				Type       string `json:"type"`
				Properties struct {
					TipoDocumento struct {
						Type        string   `json:"type"`
						Description string   `json:"description"`
						Enum        []string `json:"enum"`
					} `json:"tipoDocumento"`
					TipoGeneracion struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Enum        []int  `json:"enum"`
					} `json:"tipoGeneracion"`
					NumeroDocumento struct {
						Description string `json:"description"`
						Type        string `json:"type"`
						MinLength   int    `json:"minLength"`
						MaxLength   int    `json:"maxLength"`
					} `json:"numeroDocumento"`
					FechaEmision struct {
						Description string `json:"description"`
						Type        string `json:"type"`
						Format      string `json:"format"`
					} `json:"fechaEmision"`
				} `json:"properties"`
				AllOf []struct {
					If struct {
						Properties struct {
							TipoGeneracion struct {
								Const int `json:"const"`
							} `json:"tipoGeneracion"`
						} `json:"properties"`
					} `json:"if"`
					Then struct {
						Properties struct {
							NumeroDocumento struct {
								Pattern string `json:"pattern"`
							} `json:"numeroDocumento"`
						} `json:"properties"`
					} `json:"then"`
				} `json:"allOf"`
				AdditionalProperties bool     `json:"additionalProperties"`
				Required             []string `json:"required"`
			} `json:"items"`
			MinItems int `json:"minItems"`
			MaxItems int `json:"maxItems"`
		} `json:"documentoRelacionado"`
		Emisor struct {
			Type        string `json:"type"`
			Description string `json:"description"`
			Properties  struct {
				Nit struct {
					Description string `json:"description"`
					Type        string `json:"type"`
					Pattern     string `json:"pattern"`
					MaxLength   int    `json:"maxLength"`
				} `json:"nit"`
				Nrc struct {
					Description string `json:"description"`
					Type        string `json:"type"`
					Pattern     string `json:"pattern"`
					MinLength   int    `json:"minLength"`
					MaxLength   int    `json:"maxLength"`
				} `json:"nrc"`
				Nombre struct {
					Type        string `json:"type"`
					Description string `json:"description"`
					MaxLength   int    `json:"maxLength"`
					MinLength   int    `json:"minLength"`
				} `json:"nombre"`
				CodActividad struct {
					Type        string `json:"type"`
					Description string `json:"description"`
					Pattern     string `json:"pattern"`
					MaxLength   int    `json:"maxLength"`
					MinLength   int    `json:"minLength"`
				} `json:"codActividad"`
				DescActividad struct {
					Type        string `json:"type"`
					Description string `json:"description"`
					MaxLength   int    `json:"maxLength"`
					MinLength   int    `json:"minLength"`
				} `json:"descActividad"`
				NombreComercial struct {
					Description string   `json:"description"`
					Type        []string `json:"type"`
					MaxLength   int      `json:"maxLength"`
					MinLength   int      `json:"minLength"`
				} `json:"nombreComercial"`
				TipoEstablecimiento struct {
					Type        string   `json:"type"`
					Description string   `json:"description"`
					Enum        []string `json:"enum"`
				} `json:"tipoEstablecimiento"`
				Direccion struct {
					Type        string `json:"type"`
					Description string `json:"description"`
					Properties  struct {
						Departamento struct {
							Type        string `json:"type"`
							Description string `json:"description"`
							Pattern     string `json:"pattern"`
						} `json:"departamento"`
						Municipio struct {
							Type        string `json:"type"`
							Description string `json:"description"`
							Pattern     string `json:"pattern"`
						} `json:"municipio"`
						Complemento struct {
							Type        string `json:"type"`
							Description string `json:"description"`
							MaxLength   int    `json:"maxLength"`
							MinLength   int    `json:"minLength"`
						} `json:"complemento"`
					} `json:"properties"`
					AllOf []struct {
						If struct {
							Properties struct {
								Departamento struct {
									Const string `json:"const"`
								} `json:"departamento"`
							} `json:"properties"`
						} `json:"if"`
						Then struct {
							Properties struct {
								Municipio struct {
									Pattern string `json:"pattern"`
								} `json:"municipio"`
							} `json:"properties"`
						} `json:"then"`
					} `json:"allOf"`
					AdditionalProperties bool     `json:"additionalProperties"`
					Required             []string `json:"required"`
				} `json:"direccion"`
				Telefono struct {
					Type        string `json:"type"`
					Description string `json:"description"`
					MinLength   int    `json:"minLength"`
					MaxLength   int    `json:"maxLength"`
				} `json:"telefono"`
				Correo struct {
					Type        string `json:"type"`
					Description string `json:"description"`
					Format      string `json:"format"`
					MaxLength   int    `json:"maxLength"`
					MinLength   int    `json:"minLength"`
				} `json:"correo"`
				CodEstableMH struct {
					Description string   `json:"description"`
					Type        []string `json:"type"`
					MaxLength   int      `json:"maxLength"`
					MinLength   int      `json:"minLength"`
				} `json:"codEstableMH"`
				CodEstable struct {
					Description string   `json:"description"`
					Type        []string `json:"type"`
					MinLength   int      `json:"minLength"`
					MaxLength   int      `json:"maxLength"`
				} `json:"codEstable"`
				CodPuntoVentaMH struct {
					Description string   `json:"description"`
					Type        []string `json:"type"`
					MaxLength   int      `json:"maxLength"`
					MinLength   int      `json:"minLength"`
				} `json:"codPuntoVentaMH"`
				CodPuntoVenta struct {
					Description string   `json:"description"`
					Type        []string `json:"type"`
					MaxLength   int      `json:"maxLength"`
					MinLength   int      `json:"minLength"`
				} `json:"codPuntoVenta"`
			} `json:"properties"`
			AdditionalProperties bool     `json:"additionalProperties"`
			Required             []string `json:"required"`
		} `json:"emisor"`
		Receptor struct {
			Type        []string `json:"type"`
			Description string   `json:"description"`
			Properties  struct {
				TipoDocumento struct {
					Type        []string `json:"type"`
					Enum        []any    `json:"enum"`
					Description string   `json:"description"`
				} `json:"tipoDocumento"`
				NumDocumento struct {
					Type        []string `json:"type"`
					Description string   `json:"description"`
					MinLength   int      `json:"minLength"`
					MaxLength   int      `json:"maxLength"`
				} `json:"numDocumento"`
				Nrc struct {
					Type        []string `json:"type"`
					Description string   `json:"description"`
					Pattern     string   `json:"pattern"`
					MinLength   int      `json:"minLength"`
					MaxLength   int      `json:"maxLength"`
				} `json:"nrc"`
				Nombre struct {
					Type        []string `json:"type"`
					Description string   `json:"description"`
					MaxLength   int      `json:"maxLength"`
					MinLength   int      `json:"minLength"`
				} `json:"nombre"`
				CodActividad struct {
					Type        []string `json:"type"`
					Description string   `json:"description"`
					Pattern     string   `json:"pattern"`
					MaxLength   int      `json:"maxLength"`
					MinLength   int      `json:"minLength"`
				} `json:"codActividad"`
				DescActividad struct {
					Type        []string `json:"type"`
					Description string   `json:"description"`
					MaxLength   int      `json:"maxLength"`
					MinLength   int      `json:"minLength"`
				} `json:"descActividad"`
				Direccion struct {
					Type        []string `json:"type"`
					Description string   `json:"description"`
					Properties  struct {
						Departamento struct {
							Type        string `json:"type"`
							Description string `json:"description"`
							Pattern     string `json:"pattern"`
						} `json:"departamento"`
						Municipio struct {
							Type        string `json:"type"`
							Description string `json:"description"`
							Pattern     string `json:"pattern"`
						} `json:"municipio"`
						Complemento struct {
							Type        string `json:"type"`
							Description string `json:"description"`
							MaxLength   int    `json:"maxLength"`
							MinLength   int    `json:"minLength"`
						} `json:"complemento"`
					} `json:"properties"`
					AllOf []struct {
						If struct {
							Properties struct {
								Departamento struct {
									Const string `json:"const"`
								} `json:"departamento"`
							} `json:"properties"`
						} `json:"if"`
						Then struct {
							Properties struct {
								Municipio struct {
									Pattern string `json:"pattern"`
								} `json:"municipio"`
							} `json:"properties"`
						} `json:"then"`
					} `json:"allOf"`
					AdditionalProperties bool     `json:"additionalProperties"`
					Required             []string `json:"required"`
				} `json:"direccion"`
				Telefono struct {
					Type        []string `json:"type"`
					Description string   `json:"description"`
					MinLength   int      `json:"minLength"`
					MaxLength   int      `json:"maxLength"`
				} `json:"telefono"`
				Correo struct {
					Type        []string `json:"type"`
					Description string   `json:"description"`
					Format      string   `json:"format"`
					MaxLength   int      `json:"maxLength"`
				} `json:"correo"`
			} `json:"properties"`
			AllOf []struct {
				If struct {
					Properties struct {
						TipoDocumento struct {
							Const string `json:"const"`
						} `json:"tipoDocumento"`
					} `json:"properties"`
				} `json:"if"`
				Then struct {
					Properties struct {
						NumDocumento struct {
							Type    string `json:"type"`
							Pattern string `json:"pattern"`
						} `json:"numDocumento"`
					} `json:"properties"`
				} `json:"then"`
				Else struct {
					Properties struct {
						Nrc struct {
							Type string `json:"type"`
						} `json:"nrc"`
					} `json:"properties"`
				} `json:"else,omitempty"`
			} `json:"allOf"`
			AdditionalProperties bool     `json:"additionalProperties"`
			Required             []string `json:"required"`
		} `json:"receptor"`
		OtrosDocumentos struct {
			Description string   `json:"description"`
			Type        []string `json:"type"`
			Items       struct {
				Type       string `json:"type"`
				Properties struct {
					CodDocAsociado struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Minimum     int    `json:"minimum"`
						Maximum     int    `json:"maximum"`
					} `json:"codDocAsociado"`
					DescDocumento struct {
						Type        []string `json:"type"`
						Description string   `json:"description"`
						MaxLength   int      `json:"maxLength"`
					} `json:"descDocumento"`
					DetalleDocumento struct {
						Type        []string `json:"type"`
						Description string   `json:"description"`
						MaxLength   int      `json:"maxLength"`
					} `json:"detalleDocumento"`
					Medico struct {
						Description string   `json:"description"`
						Type        []string `json:"type"`
						Properties  struct {
							Nombre struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								MaxLength   int    `json:"maxLength"`
							} `json:"nombre"`
							Nit struct {
								Type        []string `json:"type"`
								Description string   `json:"description"`
								Pattern     string   `json:"pattern"`
							} `json:"nit"`
							DocIdentificacion struct {
								Type        []string `json:"type"`
								Description string   `json:"description"`
								MaxLength   int      `json:"maxLength"`
								MinLength   int      `json:"minLength"`
							} `json:"docIdentificacion"`
							TipoServicio struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Minimum     int    `json:"minimum"`
								Maximum     int    `json:"maximum"`
							} `json:"tipoServicio"`
						} `json:"properties"`
						AllOf []struct {
							If struct {
								Properties struct {
									Nit struct {
										Type string `json:"type"`
									} `json:"nit"`
								} `json:"properties"`
							} `json:"if"`
							Then struct {
								Properties struct {
									DocIdentificacion struct {
										Type string `json:"type"`
									} `json:"docIdentificacion"`
								} `json:"properties"`
							} `json:"then"`
						} `json:"allOf"`
						AdditionalProperties bool     `json:"additionalProperties"`
						Required             []string `json:"required"`
					} `json:"medico"`
				} `json:"properties"`
				AllOf []struct {
					If struct {
						Properties struct {
							CodDocAsociado struct {
								Const int `json:"const"`
							} `json:"codDocAsociado"`
						} `json:"properties"`
					} `json:"if"`
					Then struct {
						Properties struct {
							Medico struct {
								Type string `json:"type"`
							} `json:"medico"`
							DescDocumento struct {
								Type string `json:"type"`
							} `json:"descDocumento"`
							DetalleDocumento struct {
								Type string `json:"type"`
							} `json:"detalleDocumento"`
						} `json:"properties"`
					} `json:"then"`
					Else struct {
						Properties struct {
							DescDocumento struct {
								Type string `json:"type"`
							} `json:"descDocumento"`
							DetalleDocumento struct {
								Type string `json:"type"`
							} `json:"detalleDocumento"`
							Medico struct {
								Type string `json:"type"`
							} `json:"medico"`
						} `json:"properties"`
					} `json:"else"`
				} `json:"allOf"`
				AdditionalProperties bool     `json:"additionalProperties"`
				Required             []string `json:"required"`
			} `json:"items"`
			MinItems int `json:"minItems"`
			MaxItems int `json:"maxItems"`
		} `json:"otrosDocumentos"`
		VentaTercero struct {
			Description string   `json:"description"`
			Type        []string `json:"type"`
			Properties  struct {
				Nit struct {
					Type        string `json:"type"`
					Description string `json:"description"`
					Pattern     string `json:"pattern"`
				} `json:"nit"`
				Nombre struct {
					Type        string `json:"type"`
					Description string `json:"description"`
					MaxLength   int    `json:"maxLength"`
					MinLength   int    `json:"minLength"`
				} `json:"nombre"`
			} `json:"properties"`
			AdditionalProperties bool     `json:"additionalProperties"`
			Required             []string `json:"required"`
		} `json:"ventaTercero"`
		CuerpoDocumento struct {
			Type        string `json:"type"`
			Description string `json:"description"`
			Items       struct {
				Type       string `json:"type"`
				Properties struct {
					NumItem struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Minimum     int    `json:"minimum"`
						Maximum     int    `json:"maximum"`
					} `json:"numItem"`
					TipoItem struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Enum        []int  `json:"enum"`
					} `json:"tipoItem"`
					NumeroDocumento struct {
						Description string   `json:"description"`
						Type        []string `json:"type"`
						MinLength   int      `json:"minLength"`
						MaxLength   int      `json:"maxLength"`
					} `json:"numeroDocumento"`
					Cantidad struct {
						Type             string  `json:"type"`
						Description      string  `json:"description"`
						ExclusiveMaximum int64   `json:"exclusiveMaximum"`
						ExclusiveMinimum int     `json:"exclusiveMinimum"`
						MultipleOf       float64 `json:"multipleOf"`
					} `json:"cantidad"`
					Codigo struct {
						Description string   `json:"description"`
						Type        []string `json:"type"`
						MaxLength   int      `json:"maxLength"`
						MinLength   int      `json:"minLength"`
					} `json:"codigo"`
					CodTributo struct {
						Description string   `json:"description"`
						Type        []string `json:"type"`
						Enum        []any    `json:"enum"`
						MaxLength   int      `json:"maxLength"`
						MinLength   int      `json:"minLength"`
					} `json:"codTributo"`
					UniMedida struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Minimum     int    `json:"minimum"`
						Maximum     int    `json:"maximum"`
					} `json:"uniMedida"`
					Descripcion struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						MaxLength   int    `json:"maxLength"`
					} `json:"descripcion"`
					PrecioUni struct {
						Type             string  `json:"type"`
						Description      string  `json:"description"`
						ExclusiveMaximum int64   `json:"exclusiveMaximum"`
						MultipleOf       float64 `json:"multipleOf"`
					} `json:"precioUni"`
					MontoDescu struct {
						Type             string  `json:"type"`
						Description      string  `json:"description"`
						Minimum          int     `json:"minimum"`
						ExclusiveMaximum int64   `json:"exclusiveMaximum"`
						MultipleOf       float64 `json:"multipleOf"`
					} `json:"montoDescu"`
					VentaNoSuj struct {
						Type             string  `json:"type"`
						Description      string  `json:"description"`
						Minimum          int     `json:"minimum"`
						ExclusiveMaximum int64   `json:"exclusiveMaximum"`
						MultipleOf       float64 `json:"multipleOf"`
					} `json:"ventaNoSuj"`
					VentaExenta struct {
						Type             string  `json:"type"`
						Description      string  `json:"description"`
						Minimum          int     `json:"minimum"`
						ExclusiveMaximum int64   `json:"exclusiveMaximum"`
						MultipleOf       float64 `json:"multipleOf"`
					} `json:"ventaExenta"`
					VentaGravada struct {
						Type             string  `json:"type"`
						Description      string  `json:"description"`
						Minimum          int     `json:"minimum"`
						ExclusiveMaximum int64   `json:"exclusiveMaximum"`
						MultipleOf       float64 `json:"multipleOf"`
					} `json:"ventaGravada"`
					Tributos struct {
						Description string   `json:"description"`
						Type        []string `json:"type"`
						Items       struct {
							Type      string `json:"type"`
							MaxLength int    `json:"maxLength"`
							MinLength int    `json:"minLength"`
						} `json:"items"`
						MinItems    int  `json:"minItems"`
						UniqueItems bool `json:"uniqueItems"`
					} `json:"tributos"`
					Psv struct {
						Type             string  `json:"type"`
						Description      string  `json:"description"`
						Minimum          int     `json:"minimum"`
						ExclusiveMaximum int64   `json:"exclusiveMaximum"`
						MultipleOf       float64 `json:"multipleOf"`
					} `json:"psv"`
					NoGravado struct {
						Type             string  `json:"type"`
						Description      string  `json:"description"`
						ExclusiveMaximum int64   `json:"exclusiveMaximum"`
						ExclusiveMinimum int64   `json:"exclusiveMinimum"`
						MultipleOf       float64 `json:"multipleOf"`
					} `json:"noGravado"`
					IvaItem struct {
						Type             string  `json:"type"`
						Description      string  `json:"description"`
						Minimum          int     `json:"minimum"`
						ExclusiveMaximum int64   `json:"exclusiveMaximum"`
						MultipleOf       float64 `json:"multipleOf"`
					} `json:"ivaItem"`
				} `json:"properties"`
				AllOf []struct {
					If struct {
						Properties struct {
							VentaGravada struct {
								Maximum int `json:"maximum"`
							} `json:"ventaGravada"`
						} `json:"properties"`
					} `json:"if"`
					Then struct {
						Properties struct {
							Tributos struct {
								Type string `json:"type"`
							} `json:"tributos"`
							IvaItem struct {
								Maximum int `json:"maximum"`
							} `json:"ivaItem"`
						} `json:"properties"`
					} `json:"then"`
					Else struct {
						Properties struct {
							CodTributo struct {
								Type string `json:"type"`
							} `json:"codTributo"`
							Tributos struct {
								Items struct {
									Enum []string `json:"enum"`
								} `json:"items"`
							} `json:"tributos"`
						} `json:"properties"`
					} `json:"else,omitempty"`
				} `json:"allOf"`
				AdditionalProperties bool     `json:"additionalProperties"`
				Required             []string `json:"required"`
			} `json:"items"`
			MinItems int `json:"minItems"`
			MaxItems int `json:"maxItems"`
		} `json:"cuerpoDocumento"`
		Resumen struct {
			Type        string `json:"type"`
			Description string `json:"description"`
			Properties  struct {
				TotalNoSuj struct {
					Type             string  `json:"type"`
					Description      string  `json:"description"`
					Minimum          int     `json:"minimum"`
					ExclusiveMaximum int64   `json:"exclusiveMaximum"`
					MultipleOf       float64 `json:"multipleOf"`
				} `json:"totalNoSuj"`
				TotalExenta struct {
					Type             string  `json:"type"`
					Description      string  `json:"description"`
					Minimum          int     `json:"minimum"`
					ExclusiveMaximum int64   `json:"exclusiveMaximum"`
					MultipleOf       float64 `json:"multipleOf"`
				} `json:"totalExenta"`
				TotalGravada struct {
					Type             string  `json:"type"`
					Description      string  `json:"description"`
					Minimum          int     `json:"minimum"`
					ExclusiveMaximum int64   `json:"exclusiveMaximum"`
					MultipleOf       float64 `json:"multipleOf"`
				} `json:"totalGravada"`
				SubTotalVentas struct {
					Type             string  `json:"type"`
					Description      string  `json:"description"`
					Minimum          int     `json:"minimum"`
					ExclusiveMaximum int64   `json:"exclusiveMaximum"`
					MultipleOf       float64 `json:"multipleOf"`
				} `json:"subTotalVentas"`
				DescuNoSuj struct {
					Type             string  `json:"type"`
					Description      string  `json:"description"`
					ExclusiveMaximum int64   `json:"exclusiveMaximum"`
					Minimum          int     `json:"minimum"`
					MultipleOf       float64 `json:"multipleOf"`
				} `json:"descuNoSuj"`
				DescuExenta struct {
					Type             string  `json:"type"`
					Description      string  `json:"description"`
					ExclusiveMaximum int64   `json:"exclusiveMaximum"`
					Minimum          int     `json:"minimum"`
					MultipleOf       float64 `json:"multipleOf"`
				} `json:"descuExenta"`
				DescuGravada struct {
					Type             string  `json:"type"`
					Description      string  `json:"description"`
					ExclusiveMaximum int64   `json:"exclusiveMaximum"`
					Minimum          int     `json:"minimum"`
					MultipleOf       float64 `json:"multipleOf"`
				} `json:"descuGravada"`
				PorcentajeDescuento struct {
					Type        string  `json:"type"`
					Description string  `json:"description"`
					Maximum     int     `json:"maximum"`
					Minimum     int     `json:"minimum"`
					MultipleOf  float64 `json:"multipleOf"`
				} `json:"porcentajeDescuento"`
				TotalDescu struct {
					Type             string  `json:"type"`
					Description      string  `json:"description"`
					Minimum          int     `json:"minimum"`
					ExclusiveMaximum int64   `json:"exclusiveMaximum"`
					MultipleOf       float64 `json:"multipleOf"`
				} `json:"totalDescu"`
				Tributos struct {
					Type        []string `json:"type"`
					UniqueItems bool     `json:"uniqueItems"`
					Description string   `json:"description"`
					Items       struct {
						Type       string `json:"type"`
						Properties struct {
							Codigo struct {
								Description string   `json:"description"`
								Type        string   `json:"type"`
								MinLength   int      `json:"minLength"`
								MaxLength   int      `json:"maxLength"`
								Enum        []string `json:"enum"`
							} `json:"codigo"`
							Descripcion struct {
								Description string `json:"description"`
								Type        string `json:"type"`
								MinLength   int    `json:"minLength"`
								MaxLength   int    `json:"maxLength"`
							} `json:"descripcion"`
							Valor struct {
								Description      string  `json:"description"`
								Type             string  `json:"type"`
								Minimum          int     `json:"minimum"`
								ExclusiveMaximum int64   `json:"exclusiveMaximum"`
								MultipleOf       float64 `json:"multipleOf"`
							} `json:"valor"`
						} `json:"properties"`
						AdditionalProperties bool     `json:"additionalProperties"`
						Required             []string `json:"required"`
					} `json:"items"`
				} `json:"tributos"`
				SubTotal struct {
					Type             string  `json:"type"`
					Description      string  `json:"description"`
					Minimum          int     `json:"minimum"`
					ExclusiveMaximum int64   `json:"exclusiveMaximum"`
					MultipleOf       float64 `json:"multipleOf"`
				} `json:"subTotal"`
				IvaRete1 struct {
					Type             string  `json:"type"`
					Description      string  `json:"description"`
					Minimum          int     `json:"minimum"`
					ExclusiveMaximum int64   `json:"exclusiveMaximum"`
					MultipleOf       float64 `json:"multipleOf"`
				} `json:"ivaRete1"`
				ReteRenta struct {
					Type             string  `json:"type"`
					Description      string  `json:"description"`
					Minimum          int     `json:"minimum"`
					ExclusiveMaximum int64   `json:"exclusiveMaximum"`
					MultipleOf       float64 `json:"multipleOf"`
				} `json:"reteRenta"`
				MontoTotalOperacion struct {
					Type             string  `json:"type"`
					Description      string  `json:"description"`
					Minimum          int     `json:"minimum"`
					ExclusiveMaximum int64   `json:"exclusiveMaximum"`
					MultipleOf       float64 `json:"multipleOf"`
				} `json:"montoTotalOperacion"`
				TotalNoGravado struct {
					Type             string  `json:"type"`
					Description      string  `json:"description"`
					ExclusiveMaximum int64   `json:"exclusiveMaximum"`
					ExclusiveMinimum int64   `json:"exclusiveMinimum"`
					MultipleOf       float64 `json:"multipleOf"`
				} `json:"totalNoGravado"`
				TotalPagar struct {
					Type             string  `json:"type"`
					Description      string  `json:"description"`
					Minimum          int     `json:"minimum"`
					ExclusiveMaximum int64   `json:"exclusiveMaximum"`
					MultipleOf       float64 `json:"multipleOf"`
				} `json:"totalPagar"`
				TotalLetras struct {
					Type        string `json:"type"`
					Description string `json:"description"`
					MaxLength   int    `json:"maxLength"`
				} `json:"totalLetras"`
				TotalIva struct {
					Type             string  `json:"type"`
					Description      string  `json:"description"`
					Minimum          int     `json:"minimum"`
					ExclusiveMaximum int64   `json:"exclusiveMaximum"`
					MultipleOf       float64 `json:"multipleOf"`
				} `json:"totalIva"`
				SaldoFavor struct {
					Type             string  `json:"type"`
					Description      string  `json:"description"`
					Maximum          int     `json:"maximum"`
					ExclusiveMaximum int64   `json:"exclusiveMaximum"`
					MultipleOf       float64 `json:"multipleOf"`
				} `json:"saldoFavor"`
				CondicionOperacion struct {
					Type        string `json:"type"`
					Description string `json:"description"`
					Enum        []int  `json:"enum"`
				} `json:"condicionOperacion"`
				Pagos struct {
					Description string   `json:"description"`
					Type        []string `json:"type"`
					Items       struct {
						Type       string `json:"type"`
						Properties struct {
							Codigo struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								MaxLength   int    `json:"maxLength"`
								Pattern     string `json:"pattern"`
							} `json:"codigo"`
							MontoPago struct {
								Type             string  `json:"type"`
								Description      string  `json:"description"`
								Minimum          int     `json:"minimum"`
								ExclusiveMaximum int64   `json:"exclusiveMaximum"`
								MultipleOf       float64 `json:"multipleOf"`
							} `json:"montoPago"`
							Referencia struct {
								Type        []string `json:"type"`
								Description string   `json:"description"`
								MaxLength   int      `json:"maxLength"`
							} `json:"referencia"`
							Plazo struct {
								Description string   `json:"description"`
								Type        []string `json:"type"`
								Pattern     string   `json:"pattern"`
							} `json:"plazo"`
							Periodo struct {
								Description string   `json:"description"`
								Type        []string `json:"type"`
							} `json:"periodo"`
						} `json:"properties"`
						AdditionalProperties bool     `json:"additionalProperties"`
						Required             []string `json:"required"`
					} `json:"items"`
					MinItems int `json:"minItems"`
				} `json:"pagos"`
				NumPagoElectronico struct {
					Description string   `json:"description"`
					Type        []string `json:"type"`
					MaxLength   int      `json:"maxLength"`
				} `json:"numPagoElectronico"`
			} `json:"properties"`
			AllOf []struct {
				If struct {
					Properties struct {
						TotalGravada struct {
							Maximum int `json:"maximum"`
						} `json:"totalGravada"`
					} `json:"properties"`
				} `json:"if"`
				Then struct {
					Properties struct {
						IvaRete1 struct {
							Maximum int `json:"maximum"`
						} `json:"ivaRete1"`
					} `json:"properties"`
				} `json:"then"`
			} `json:"allOf"`
			AdditionalProperties bool     `json:"additionalProperties"`
			Required             []string `json:"required"`
		} `json:"resumen"`
		Extension struct {
			Type        []string `json:"type"`
			Description string   `json:"description"`
			Properties  struct {
				NombEntrega struct {
					Description string   `json:"description"`
					Type        []string `json:"type"`
					MaxLength   int      `json:"maxLength"`
					MinLength   int      `json:"minLength"`
				} `json:"nombEntrega"`
				DocuEntrega struct {
					Description string   `json:"description"`
					Type        []string `json:"type"`
					MaxLength   int      `json:"maxLength"`
					MinLength   int      `json:"minLength"`
				} `json:"docuEntrega"`
				NombRecibe struct {
					Description string   `json:"description"`
					Type        []string `json:"type"`
					MaxLength   int      `json:"maxLength"`
					MinLength   int      `json:"minLength"`
				} `json:"nombRecibe"`
				DocuRecibe struct {
					Description string   `json:"description"`
					Type        []string `json:"type"`
					MaxLength   int      `json:"maxLength"`
					MinLength   int      `json:"minLength"`
				} `json:"docuRecibe"`
				Observaciones struct {
					Description string   `json:"description"`
					Type        []string `json:"type"`
					MaxLength   int      `json:"maxLength"`
				} `json:"observaciones"`
				PlacaVehiculo struct {
					Description string   `json:"description"`
					Type        []string `json:"type"`
					MaxLength   int      `json:"maxLength"`
				} `json:"placaVehiculo"`
			} `json:"properties"`
			AdditionalProperties bool     `json:"additionalProperties"`
			Required             []string `json:"required"`
		} `json:"extension"`
		Apendice struct {
			Description string   `json:"description"`
			Type        []string `json:"type"`
			Items       struct {
				Type       string `json:"type"`
				Properties struct {
					Campo struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						MaxLength   int    `json:"maxLength"`
					} `json:"campo"`
					Etiqueta struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						MaxLength   int    `json:"maxLength"`
					} `json:"etiqueta"`
					Valor struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						MaxLength   int    `json:"maxLength"`
					} `json:"valor"`
				} `json:"properties"`
				AdditionalProperties bool     `json:"additionalProperties"`
				Required             []string `json:"required"`
			} `json:"items"`
			MinItems int `json:"minItems"`
			MaxItems int `json:"maxItems"`
		} `json:"apendice"`
	} `json:"properties"`
	AdditionalProperties bool `json:"additionalProperties"`
	AllOf                []struct {
		If struct {
			Properties struct {
				Resumen struct {
					Properties struct {
						MontoTotalOperacion struct {
							Minimum float64 `json:"minimum"`
						} `json:"montoTotalOperacion"`
					} `json:"properties"`
				} `json:"resumen"`
			} `json:"properties"`
		} `json:"if"`
		Then struct {
			Properties struct {
				Receptor struct {
					Type       string `json:"type"`
					Properties struct {
						TipoDocumento struct {
							Type string `json:"type"`
						} `json:"tipoDocumento"`
						NumDocumento struct {
							Type string `json:"type"`
						} `json:"numDocumento"`
						Nombre struct {
							Type string `json:"type"`
						} `json:"nombre"`
					} `json:"properties"`
				} `json:"receptor"`
			} `json:"properties"`
		} `json:"then"`
	} `json:"allOf"`
	Required []string `json:"required"`
}
